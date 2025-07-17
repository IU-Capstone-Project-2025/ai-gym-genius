interface UserStore {
    token: Ref<string>,
    login: (username: string, password: string) => Promise<void>
}

export const useUserStore = defineStore('userStore', (): UserStore => {
    const token = useCookie('fucking-auth-token');
    const apiURI = 'http://api.говно.site/auth_admin'

    function login(username: string, password: string) {
        return $fetch(apiURI, {
            method: 'POST',
            body: {
                login: username,
                password: password
            }
        })
    }

    return {
        token,
        login
    }
})