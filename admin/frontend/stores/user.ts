interface UserStore {
    token: Ref<string>,
    login: (username: string, password: string) => Promise<void>,
    isAuthorized: ComputedRef<boolean>,
    logout: () => void
}

export const useUserStore = defineStore('userStore', (): UserStore => {
    const token = useCookie('fucking-auth-token');
    const apiURI = 'http://api.говно.site/auth_admin'

    function login(username: string, password: string) {
        // todo: actually fetch api token from the backend (now it`s impossible since backend restricted to extract response headers)
        token.value = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTMwMTYzNzEsImlkIjoxLCJsb2dpbiI6ImFkbWluIiwicm9sZSI6ImFkbWluIn0.HQk3ZZggYSZB9eSdV1yB00A-BwNYxOuFK_Kr24Lk0U0";

        return $fetch(apiURI, {
            method: 'POST',
            body: {
                login: username,
                password: password
            }
        })
    }

    function logout() {
        token.value = "";
        navigateTo('/auth');
    }

    const isAuthorized = computed(() => {
        console.log('TOKEN: ', token.value);
        return !!token.value;
    })

    return {
        token,
        login,
        logout,
        isAuthorized
    }
})