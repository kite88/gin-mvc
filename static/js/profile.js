class Profile {
    key = 'userInfo';

    save(profile) {
        sessionStorage.setItem(this.key, JSON.stringify(profile))
    }

    get() {
        return JSON.parse(sessionStorage.getItem(this.key))
    }

    del() {
        sessionStorage.removeItem(this.key)
    }

    getUser() {
        var data = this.get()
        if (data) {
            return data.user
        }
        return null
    }

    getToken() {
        var data = this.get()
        if (data) {
            return data.Token.token
        }
        return ""
    }
}
