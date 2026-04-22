import { ref, type Ref } from 'vue'
import type { CurrentStepRequestMessage, ExpeditionStepResolveInfo, Hero, User, Waifu } from './types';
import type { VueCookies } from 'vue-cookies';

class global {
    public static readonly DOMAIN_NAME = "https://tama.rhiobet.sh";

    //nanapi 
    public static readonly NANAPI_DOMAIN = "https://https://waicolle.japan7.bde.enseeiht.fr/"

    //Request object by id
    public static readonly REQ_HERO = "/hero/{id}";
    public static readonly REQ_TEAM = "/teams/{id}";
    public static readonly REQ_GAMESTATE = "/gamestate/{id}";
    public static readonly REQ_USR = "/user/{id}";
    public static readonly REQ_AVAILABLEEXPEDITIONS = "/availableexpeditions/"
    public static readonly REQ_CURRENTEXPEDITIONSTEP = "/currentexpeditionstep/";

    //request update objects
    public static readonly REQ_LAUNCHEXPEDITION = "/launchExpedition/{usr}/{expId}";
    public static readonly REQ_UPDATETEAM = "/updateTeam/";

    //Create objects
    public static readonly REQ_CREATEHEROFROMWAIFU = "/createherofromwaifu/{id}"

    //nanapi requests 
    public static readonly REQ_USERWAIFUS = "/userwaifus/{id}"

    public static readonly NO_IMAGE = "/imgs/noimage.jpg";
    public static readonly EXPEDITION = "/imgs/expedition.png";
}

enum HomeStatus {
    Noting = 1,
    ExpeditionManagement,
    TeamManagement,
    HeroGetter,
    InspectHero,
}

class NavigationHandler {
    homeStatus = ref(HomeStatus.Noting)
    heroToInspect = ref<Hero | undefined>(undefined)

    constructor() {
        this.homeStatus.value = HomeStatus.Noting
    }

    setHomeStatus(newHomeStatus: HomeStatus): void {
        this.homeStatus.value = newHomeStatus
    }

    getHomeStatus() {
        return this.homeStatus
    }

    setHeroToInspect(h: Hero) {
        this.heroToInspect.value = h
    }

    getHeroToInspect() {
        return this.heroToInspect
    }
}

enum RequestType {
    Hero = 1,
    Team,
    GameState,
    User,
    AvailableExpeditions,
    CurrentExpeditionStep,
    UserWaifus,
    CreateHeroFromWaifu,

    LaunchExpedition,
    UpdateTeam,
}

function buildRequestPath(reqType: RequestType, pathParams: { id: string; value: string }[] | undefined = undefined): string {
    let request: string = global.DOMAIN_NAME
    switch (reqType) {
        case RequestType.Hero:
            request += global.REQ_HERO
            break
        case RequestType.Team:
            request += global.REQ_TEAM
            break
        case RequestType.GameState:
            request += global.REQ_GAMESTATE
            break
        case RequestType.User:
            request += global.REQ_USR
            break
        case RequestType.AvailableExpeditions:
            request += global.REQ_AVAILABLEEXPEDITIONS
            break
        case RequestType.CurrentExpeditionStep:
            request += global.REQ_CURRENTEXPEDITIONSTEP
            break
        case RequestType.UserWaifus:
            request += global.REQ_USERWAIFUS
            break
        case RequestType.CreateHeroFromWaifu:
            request += global.REQ_CREATEHEROFROMWAIFU
            break
        case RequestType.LaunchExpedition:
            request += global.REQ_LAUNCHEXPEDITION
            break
        case RequestType.UpdateTeam:
            request += global.REQ_UPDATETEAM
            break
        default:
            request = ""
            break
    }

    if (pathParams) {
        for (const pathParam of pathParams) {
            request = request.replace(`{${pathParam.id}}`, pathParam.value)
        }
    }

    return request
}

async function fetchData<T>(target: Ref<T | undefined>, reqType: RequestType, pathParams: [{ id: string; value: string }] | undefined = undefined) {
    target.value = undefined;

    let request: string = buildRequestPath(reqType, pathParams)
    console.log(request)
    if (request !== "") {
        console.log("sending get request to : " + request)
        const res = await fetch(request)
        target.value = await res.json() as T
    }
}

async function postRequest<AnswerType, BodyType>(
    answer: Ref<AnswerType | undefined>,
    dataToSend: BodyType,
    requestType: RequestType,
    pathParams: [{ id: string; value: string }] | undefined = undefined) {

    answer.value = undefined;
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(dataToSend)
    };

    let request = buildRequestPath(requestType, pathParams)
    console.log(request)
    if (request !== "") {
        console.log("sending post request to : " + request)
        const res = await fetch(request, requestOptions)
        answer.value = await res.json() as AnswerType
    }
}

async function getCurrentExpeditionStepResolveInfo(answer: Ref<ExpeditionStepResolveInfo | undefined>, usreid: number) {
    let message: CurrentStepRequestMessage = {
        id: usreid,
        time: Date.now()
    }
    await postRequest<ExpeditionStepResolveInfo, CurrentStepRequestMessage>(answer, message, RequestType.CurrentExpeditionStep)
}

async function launchExpedition(target: Ref<ExpeditionStepResolveInfo | undefined>, user: User, expIdentifier: string) {
    target.value = undefined
    let request: string = buildRequestPath(RequestType.LaunchExpedition)
    request = request.replace(`{usr}`, String(user.id))
    request = request.replace(`{expId}`, expIdentifier)
    const response = await fetch(request);
    target.value = await response.json() as ExpeditionStepResolveInfo
    if (target.value) {
        user.state.state = target.value.stepState
    }
}

async function createAHeroFromAWaifu(target: Ref<Hero | undefined>, waifu: Waifu, user: User) {
    console.log(waifu)
    postRequest<Hero, Waifu>(target, waifu, RequestType.CreateHeroFromWaifu, [{ id: "id", value: `${user.id}` }])
}

function formatTextTimeFromTimeMS(timeMS: number) {
    let res = ""
    //console.log(distance)
    // Time calculations for days, hours, minutes and seconds
    var days = Math.floor(timeMS / (1000 * 60 * 60 * 24));
    var hours = Math.floor((timeMS % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
    var minutes = Math.floor((timeMS % (1000 * 60 * 60)) / (1000 * 60));
    var seconds = Math.floor((timeMS % (1000 * 60)) / 1000);

    if (seconds > 0) {
        res = seconds + "s"
    }
    if (minutes > 0) {
        res = minutes + "m " + res
    }
    if (hours > 0) {
        res = hours + "h " + res
    }
    if (days > 0) {
        res = days + "d " + res
    }
    return res
}

function isUserIdInRequestParams(): boolean {
    let urlParams = new URLSearchParams(window.location.search);
    return urlParams.has('wvtrusrid')
}

function isUserIdInCookies($cookies: VueCookies): boolean {
    return $cookies.get("wvtrusrid") != undefined
}

function getUserIdFromRequestParams(): string | null {
    let urlParams = new URLSearchParams(window.location.search);
    let res = urlParams.get('wvtrusrid')
    return !res ? null : res
}

function getUserIdFromCookies($cookies: VueCookies): string | null {
    return $cookies.get("wvtrusrid")
}

function getUserIDFromCookiesOrURLParams($cookies: VueCookies | undefined) {
    let wvtrusrid: string | null = null
    let urlParams = new URLSearchParams(window.location.search);
    if ($cookies != undefined && $cookies.get("wvtrusrid")) {
        wvtrusrid = $cookies.get("wvtrusrid")
        console.log("cookies uid : " + wvtrusrid)
    } else if (urlParams.has('wvtrusrid')) {
        wvtrusrid = urlParams.get('wvtrusrid')
        console.log("url param uid : " + wvtrusrid)
    }
    return wvtrusrid
}

export {
    global,
    fetchData,
    postRequest,
    launchExpedition,
    getCurrentExpeditionStepResolveInfo,
    formatTextTimeFromTimeMS,
    getUserIDFromCookiesOrURLParams,
    createAHeroFromAWaifu,
    isUserIdInCookies,
    isUserIdInRequestParams,
    getUserIdFromRequestParams,
    getUserIdFromCookies,
    RequestType,
    HomeStatus,
    NavigationHandler,
}
