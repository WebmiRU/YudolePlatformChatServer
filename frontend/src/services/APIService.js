var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
class APIService {
    modulesIndexGet() {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield fetch('/api/modules');
            return yield response.json();
        });
    }
    modulesIdGet(id) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield fetch('/api/modules/' + id);
            return yield response.json();
        });
    }
    modulesIdPut(id, payload) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield fetch('/api/modules/' + id, {
                headers: {
                    'Content-type': 'application/json'
                },
                method: 'PUT',
                body: JSON.stringify(payload),
            });
            return yield response.json();
        });
    }
    putModulesIdSetState(id, state) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield fetch('/api/modules/' + id + '/state/' + state, {
                method: 'PUT',
            });
            return yield response.json();
        });
    }
    modulesIdStart(id) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield fetch('/api/modules/' + id + '/start', {
                method: 'POST',
            });
            return yield response.json();
        });
    }
    modulesIdStop(id) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield fetch('/api/modules/' + id + '/stop', {
                method: 'POST',
            });
            return yield response.json();
        });
    }
    themesIndex() {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield fetch('/api/themes', {
                method: 'GET',
            });
            return yield response.json();
        });
    }
    themesGet(id) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield fetch('/api/themes/' + id, {
                method: 'GET',
            });
            return yield response.json();
        });
    }
    themesPut(id, payload) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield fetch('/api/themes/' + id, {
                headers: {
                    'Content-type': 'application/json'
                },
                method: 'PUT',
                body: JSON.stringify(payload),
            });
            return yield response.json();
        });
    }
    channelsIndex() {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield fetch('/api/channels', {
                method: 'GET',
            });
            return yield response.json();
        });
    }
    channelsGet(id) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield fetch('/api/channels/' + id, {
                method: 'GET',
            });
            return yield response.json();
        });
    }
    channelsPut(id, payload) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield fetch('/api/channels/' + id, {
                headers: {
                    'Content-type': 'application/json'
                },
                method: 'PUT',
                body: JSON.stringify(payload),
            });
            return yield response.json();
        });
    }
    apiPut(url, model) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield fetch(url, {
                headers: {
                    'Content-type': 'application/json'
                },
                method: 'PUT',
                body: JSON.stringify(model),
            });
            return yield response.json();
        });
    }
    apiGet(url) {
        return __awaiter(this, void 0, void 0, function* () {
            const response = yield fetch(url, {
                method: 'GET',
            });
            return yield response.json();
        });
    }
}
export default new APIService;
