import type { Ref } from 'vue'

class global {
    public static readonly DOMAIN_NAME = "https://tama.rhiobet.sh";
    public static readonly REQ_WAIFU = "/waifus/{id}";
    public static readonly REQ_TEAM = "/teams/{id}";
    public static readonly REQ_GAMESTATE = "/gamestate/{id}";
    public static readonly REQ_USR = "/user/{id}";
}

enum RequestType {
    Waifu = 1,
    Team,
    GameState,
    User,
}

// async function fetchData(request: string) {
//     const res = await fetch(request)
//     return res.json()
// }

// async function fetchData<T>(reqType: RequestType, pathParams: [{ id: string; value: string }]): Promise<T | undefined> {
//     let request: string = ""
//     switch (reqType) {
//         case RequestType.Waifu:
//             request = global.DOMAIN_NAME + global.REQ_WAIFU
//         case RequestType.Team:
//             request = global.DOMAIN_NAME + global.REQ_TEAM
//         case RequestType.GameState:
//             request = global.DOMAIN_NAME + global.REQ_GAMESTATE
//         case RequestType.User:
//             request = global.DOMAIN_NAME + global.REQ_USR
//     }

//     for (const pathParam of pathParams) {
//         request = request.replace(`{${pathParam.id}}`, pathParam.value)
//     }

//     if (request !== "") {
//         const res = await fetch(request)
//         return await res.json() as T
//     }
//     return undefined
// }

async function fetchData<T>(target: Ref<T | undefined>, reqType: RequestType, pathParams: [{ id: string; value: string }]) {
    target.value = undefined;

    let request: string = ""
    switch (reqType) {
        case RequestType.Waifu:
            request = global.DOMAIN_NAME + global.REQ_WAIFU
        case RequestType.Team:
            request = global.DOMAIN_NAME + global.REQ_TEAM
        case RequestType.GameState:
            request = global.DOMAIN_NAME + global.REQ_GAMESTATE
        case RequestType.User:
            request = global.DOMAIN_NAME + global.REQ_USR
    }

    for (const pathParam of pathParams) {
        request = request.replace(`{${pathParam.id}}`, pathParam.value)
    }

    if (request !== "") {
        const res = await fetch(request)
        target.value = await res.json() as T
    }
}

export {
    global,
    fetchData,
    RequestType,
}
