import axios from 'axios'

const axiosInstance = axios.create({
    baseURL: import.meta.env.VITE_SERVER_HOST,
})

axiosInstance.interceptors.request.use(request => {
    const accessToken = localStorage.getItem('access_token');
    if (accessToken) {
        request.headers['Authorization'] = `Bearer ${accessToken}`;
    }
    return request;
}, error => {
    return Promise.reject(error);
});

axiosInstance.interceptors.response.use(
    (response) => response,
    async (error) => {
        const originalRequest = error.config;
        if (error.response.status === 401 && !originalRequest._retry) {
            try {
                originalRequest._retry = true;
                const refreshToken = localStorage.getItem('refresh_token'); // Retrieve the stored refresh token.
                // Make a request to your auth server to refresh the token.
                const response = await axios.post(import.meta.env.VITE_SERVER_HOST + '/api/v1/auth/token/refresh', {
                    refresh_token: refreshToken
                });
                const {access_token} = response.data;
                // Store the new access and refresh tokens.
                localStorage.setItem('access_token', access_token);
                // localStorage.setItem('refresh_token', newRefreshToken);
                // Update the authorization header with the new access token.
                axiosInstance.defaults.headers.common['Authorization'] = `Bearer ${access_token}`;
                return axiosInstance(originalRequest); // Retry the original request with the new access token.
            } catch (refreshError) {
                // Handle refresh token errors by clearing stored tokens and redirecting to the login page.
                console.error('Token refresh failed:', refreshError);
                localStorage.removeItem('access_token');
                localStorage.removeItem('refresh_token');
                window.location.href = '/login';
                return Promise.reject(refreshError);
            }
        }
    }
)

export default axiosInstance