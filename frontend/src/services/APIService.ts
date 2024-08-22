class APIService {
    async modulesIndexGet() {
        const response = await fetch('/api/modules')
        return await response.json();
    }

    async modulesIdGet(id: string) {
        const response = await fetch('/api/modules/' + id)
        return await response.json();
    }

    async modulesIdPut(id: string, payload: object) {
        const response = await fetch('/api/modules/' + id, {
            headers: {
                'Content-type': 'application/json'
            },
            method: 'PUT',
            body: JSON.stringify(payload),
        })
        return await response.json();
    }

    async putModulesIdSetState(id: string, state: string) {
        const response = await fetch('/api/modules/' + id + '/state/' + state, {
            method: 'PUT',
        })
        return await response.json();
    }

    async modulesIdStart(id: string) {
        const response = await fetch('/api/modules/' + id + '/start', {
            method: 'POST',
        })
        return await response.json();
    }

    async modulesIdStop(id: string) {
        const response = await fetch('/api/modules/' + id + '/stop', {
            method: 'POST',
        })
        return await response.json();
    }

    async themesIndex() {
        const response = await fetch('/api/themes', {
            method: 'GET',
        })
        return await response.json();
    }

    async themesGet(id: string) {
        const response = await fetch('/api/themes/' + id, {
            method: 'GET',
        })
        return await response.json();
    }

    async themesPut(id: string, payload: object) {
        const response = await fetch('/api/themes/' + id, {
            headers: {
                'Content-type': 'application/json'
            },
            method: 'PUT',
            body: JSON.stringify(payload),
        })
        return await response.json();
    }

    async channelsIndex() {
        const response = await fetch('/api/channels', {
            method: 'GET',
        })
        return await response.json();
    }

    async channelsGet(id: string) {
        const response = await fetch('/api/channels/' + id, {
            method: 'GET',
        })
        return await response.json();
    }

    async channelsPut(id: string, payload: object) {
        const response = await fetch('/api/channels/' + id, {
            headers: {
                'Content-type': 'application/json'
            },
            method: 'PUT',
            body: JSON.stringify(payload),
        })
        return await response.json();
    }

    async apiPut(url: string, model: object) {
        const response = await fetch(url, {
            headers: {
                'Content-type': 'application/json'
            },
            method: 'PUT',
            body: JSON.stringify(model),
        })
        return await response.json();
    }

    async apiGet(url: string) {
        const response = await fetch(url, {
            method: 'GET',
        })

        return await response.json();
    }

    async resourcesAudioGet() {
        const response = await fetch('/api/resources/audio', {
            method: 'GET',
        })
        return await response.json();
    }
}

export default new APIService
