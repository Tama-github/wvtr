import { inject, ref, type Ref } from "vue"
import { EncounterState, EquipmentType, InventoryToDo, type Armor, type CurrentStepRequestMessage, type Equipable, type EquipmentToSpend, type ExpeditionDB, type ExpeditionStepResolveInfo, type ExpeditionStepTimestamp, type ExpToLaunch, type GameState, type Hero, type Inventory, type Omamori, type Storable, type Team, type User, type Waifu, type Weapon } from "./types"
import type { VueCookies } from "vue-cookies";
import { buildRequestPath, fetchData, global, postRequest, RequestType } from "./utils";
import { buildExpeditionsCathegory, type ExpeditionCategory, type ExpToGetFromBack } from "./expeditions";

enum NavigationStatus {
    Connexion = 1,
    Home,
    ExpeditionManagement,
    TeamManagement,
    HeroMaker,
    InspectHero,
    Inventory,
}

// connexion handling 
const cookieKey = 'wvtrusrid'

function isUserIdInRequestParams(): boolean {
    let urlParams = new URLSearchParams(window.location.search);
    return urlParams.has(cookieKey)
}

function isUserIdInCookies($cookies: VueCookies): boolean {
    return $cookies.get("wvtrusrid") != undefined
}

function getUserIdFromRequestParams(): string | null {
    let urlParams = new URLSearchParams(window.location.search);
    let res = urlParams.get(cookieKey)
    return !res ? null : res
}

function getUserIdFromCookies($cookies: VueCookies): string | null {
    console.log($cookies)
    return $cookies.get(cookieKey)
}

function getUserIDFromCookiesOrURLParams($cookies: VueCookies | undefined) {
    let wvtrusrid: string | null = null
    let urlParams = new URLSearchParams(window.location.search);
    if ($cookies != undefined && $cookies.get(cookieKey)) {
        wvtrusrid = $cookies.get(cookieKey)
        console.log("cookies uid : " + wvtrusrid)
    } else if (urlParams.has(cookieKey)) {
        wvtrusrid = urlParams.get(cookieKey)
        console.log("url param uid : " + wvtrusrid)
    }
    return wvtrusrid
}

async function setupCookieWithUserID($cookies: VueCookies, uid: number) {
    console.log("user.id = " + uid)
    // we got here only if cookies have a wvtrusrid and we fetched a user
    let uidfcookie = getUserIdFromCookies($cookies!)
    if (uidfcookie == uid + "") { // check integrity of the cookie and the user
        // all good
    } else {
        // client think it has the right uid but it is wrong auth again
        requestToAuth($cookies)
    }
}

async function requestToAuth($cookies: VueCookies) {
    window.location.replace(global.REQ_AUTH);
}

class NavigationHandler {
    // page status management
    navigationStatus = ref(NavigationStatus.Connexion)
    heroToInspect = ref<Hero | undefined>(undefined)

    // the hendler is in charge of keeping it updated
    user = ref<User | undefined>(undefined)
    availableExpedition = ref<ExpeditionCategory[] | undefined>(undefined)
    //availableExpedition = ref<ExpToGetFromBack[] | undefined>(undefined)
    userWaifus = ref<Waifu[] | undefined>(undefined)
    currentExpeditionStepResolveInfo = ref<ExpeditionStepResolveInfo | undefined>(undefined)

    // Equip hero
    heroToEquip = ref<Hero | undefined>(undefined)
    expToLaunch = ref<ExpToLaunch | undefined>(undefined)
    inventoryType = ref<EquipmentType | undefined>(undefined)
    whatInventoryDo = ref<InventoryToDo | undefined>(undefined)

    constructor() {
        this.navigationStatus.value = NavigationStatus.Connexion
    }

    async closeReport() {
        await this.setGameState(EncounterState.Home)
        await this.clearCurrentExpInv()
        await this.fetchTeam()
        await this.fetchInventory()
    }

    async fetchAvailableExpedition() {
        let exps = ref<ExpToGetFromBack[] | undefined>(undefined)
        await fetchData<ExpToGetFromBack[]>(exps, RequestType.AvailableExpeditions, [{ id: "id", value: `${this.user.value!.id}` }])
        this.availableExpedition.value = buildExpeditionsCathegory(exps.value!)
    }

    async fetchCurrentExpeditionStepResolveInfo(usreid: number) {
        let message: CurrentStepRequestMessage = {
            id: usreid,
            time: Date.now()
        }
        await postRequest<ExpeditionStepResolveInfo, CurrentStepRequestMessage>(this.currentExpeditionStepResolveInfo, message, RequestType.CurrentExpeditionStep)
    }

    async fetchUserWaifus() {
        await fetchData<Waifu[]>(this.userWaifus, RequestType.UserWaifus, [{ id: "id", value: `${this.user.value!.id}` }])
        return this.userWaifus
    }

    async fetchTeam() {
        let response = ref<Team | undefined>(undefined)
        await fetchData<Team>(response, RequestType.Team, [{ id: "id", value: `${this.user.value!.currentTeam!.id}` }])

        this.user.value!.currentTeam! = response.value!
    }

    async clearCurrentExpInv() {
        let response = ref<undefined>(undefined)
        await fetchData<undefined>(response, RequestType.ClearCurrentExp, [{ id: "uid", value: `${this.user.value!.id}` }])
    }

    async fetchInventory() {
        let response = ref<Inventory | undefined>(undefined)
        await fetchData<Inventory>(response, RequestType.Inventory, [{ id: "id", value: `${this.user.value!.currentTeam!.id}` }])

        this.user.value!.inventory = response.value!
        await this.fetchAvailableExpedition()
    }

    async fetchExpeditionReport() {
        const report = ref<ExpeditionDB | undefined>(undefined)
        await fetchData<ExpeditionDB>(report, RequestType.ExpeditionReport, [{ id: "uid", value: `${this.user.value!.id}` }])
        this.user.value!.state!.currentExpedition = report.value!
    }

    getHeroToEquip() {
        return this.heroToEquip
    }

    getInventoryType() {
        return this.inventoryType
    }

    getInventoryDo() {
        return this.whatInventoryDo
    }

    getAvailableExpedition() {
        return this.availableExpedition
    }

    getUserWaifus() {
        return this.userWaifus
    }

    getReport() {
        return this.user.value?.state!.currentExpedition!
    }

    getCurrentExpeditionStepResolveInfo() {
        return this.currentExpeditionStepResolveInfo
    }

    async launchExpedition(target: Ref<ExpeditionStepResolveInfo | undefined>, expCategory: string, expIdentifier: string, giveToExpedition: EquipmentToSpend[] | undefined) {
        target.value = undefined
        let request: string = buildRequestPath(RequestType.LaunchExpedition)
        request = request.replace(`{usr}`, String(this.user.value!.id))
        request = request.replace(`{expCat}`, expCategory)
        request = request.replace(`{expId}`, expIdentifier)

        const response = await fetch(request, {
            method: "POST",
            body: JSON.stringify(giveToExpedition)
        });
        await this.fetchInventory()
        target.value = await response.json() as ExpeditionStepResolveInfo
        if (target.value) {
            this.user.value!.state!.state = target.value.stepState
        }
    }

    async createAHeroFromAWaifu(waifu: Waifu) {
        const target = ref<Hero | undefined>(undefined)
        await postRequest<Hero, Waifu>(target, waifu, RequestType.CreateHeroFromWaifu, [{ id: "id", value: `${this.user.value!.id}` }])
        this.user.value!.ownedHeroes.push(target.value!)
        this.setHeroToInspect(target.value!)
        this.setHomeStatus(NavigationStatus.InspectHero)
    }

    setHomeStatus(newHomeStatus: NavigationStatus): void {
        this.navigationStatus.value = newHomeStatus
    }

    getNavigationStatus() {
        return this.navigationStatus
    }

    setHeroToInspect(h: Hero) {
        this.heroToInspect.value = h
    }

    getExpToLaunch() {
        return this.expToLaunch
    }

    getHeroToInspect() {
        return this.heroToInspect
    }

    getUser() {
        return this.user
    }

    getUserState() {
        if (this.user.value == undefined) {
            return EncounterState.Error
        }
        return this.user.value.state!.state
    }

    async setGameState(state: EncounterState) {
        this.user.value!.state!.state = state
        this.setHomeStatus(NavigationStatus.Home)
        const answer = ref<GameState | undefined>(undefined)
        await postRequest<GameState, GameState>(answer, this.user.value!.state!, RequestType.SaveGameState)
    }

    getUserCurrentTeam() {
        return this.user.value!.currentTeam
    }

    setInventoryVue(type: EquipmentType, heroToEquip: Hero) {
        this.inventoryType.value = type
        this.heroToEquip.value = heroToEquip
        this.expToLaunch.value = undefined
        this.whatInventoryDo.value = InventoryToDo.Equip
        this.navigationStatus.value = NavigationStatus.Inventory
    }

    setInventoryForExpeditionLaunch(type: EquipmentType, exp: ExpToLaunch) {
        this.inventoryType.value = type
        this.heroToEquip.value = undefined
        this.expToLaunch.value = exp
        this.whatInventoryDo.value = InventoryToDo.Spend
        this.navigationStatus.value = NavigationStatus.Inventory
    }

    async setUserCurrentTeam(heroes: Hero[]) {
        this.user.value!.currentTeam!.heroes = heroes
        let tmpTeam = ref<Team | undefined>(undefined)
        await postRequest<Team, User>(tmpTeam, this.user.value!, RequestType.UpdateTeam)
        if (tmpTeam.value) {
            this.user.value!.currentTeam = tmpTeam.value
        }
    }

    /**
     * Setup the navigation handler with the cookies that contains the user id informations
     * @param $cookies 
     * @returns 
     */
    async setup($cookies: VueCookies) {
        if ($cookies && isUserIdInRequestParams()) { // we got here only after authantification has been done
            let uidstring = getUserIdFromRequestParams()
            $cookies.set(cookieKey, uidstring, '30d', undefined, undefined, true, "Strict")

            // redirect to the main page 
            window.location.replace("/");
        } else if ($cookies && isUserIdInCookies($cookies)) { // client has been connected once, but we need to check if it matches the database
            await fetchData<User>(this.user, RequestType.User, [{ id: "id", value: `${getUserIdFromCookies($cookies)}` }])
            if (this.user.value != undefined) {
                setupCookieWithUserID($cookies, this.user.value.id)
                this.navigationStatus.value = NavigationStatus.Home
                await this.fetchAvailableExpedition()
                await this.fetchUserWaifus()
            }
        } else { // we don't know the user and need to ask auth for it
            await requestToAuth($cookies)
        }
        return this.user
    }

    /**
     * Update the state of user team and oponent team state during expedition
     * @param user 
     * @param eri 
     * @param now 
     * @returns 
     */
    applyTimelineEventToTeam(user: Ref<User | undefined>, eri: Ref<ExpeditionStepResolveInfo | undefined>, now: number) {
        if (!user.value?.currentTeam || !eri.value?.timeline || !eri.value.eTeam) {
            return
        }

        let pvAccumulatorTeam = new Array<number>(3).fill(0)
        let pvAccumulatorEn = new Array<number>(3).fill(0)
        let index = this.getIdxOfTimeline(eri.value.timeline, now)
        for (let i = 0; i < index; i++) {
            if (!(eri.value.timeline[i]) || !eri.value.timeline[i]?.whatAction) {
                continue;
            }

            let from = eri.value.timeline[i]!.whatAction!.fromH!
            let fpvc = eri.value.timeline[i]!.whatAction!.fromPVChange!
            if (fpvc != 0) {
                let i1 = this.findHeroIdx(user.value.currentTeam, from)
                let i2 = this.findHeroIdx(eri.value.eTeam!, from)
                i1 >= 0 ? pvAccumulatorTeam[i1]! -= fpvc : (i2 >= 0 ? pvAccumulatorEn[i2]! -= fpvc : 0)
            }

            let target = eri.value.timeline[i]!.whatAction!.targetH!
            let tpvc = eri.value.timeline[i]!.whatAction!.targetPVChange!
            if (tpvc != 0) {
                let j1 = this.findHeroIdx(user.value.currentTeam, target)
                let j2 = this.findHeroIdx(eri.value.eTeam!, target)
                j1 >= 0 ? pvAccumulatorTeam[j1]! -= tpvc : (j2 >= 0 ? pvAccumulatorEn[j2]! -= tpvc : 0)
            }
        }

        for (let i = 0; i < user.value.currentTeam.heroes.length; i++) {
            user.value.currentTeam.heroes[i]!.attributes!.currentHP! = user.value.currentTeam.heroes[i]?.attributes!.maxHP! + pvAccumulatorTeam[i]!
        }

        for (let i = 0; i < eri.value.eTeam!.heroes.length; i++) {
            eri.value.eTeam.heroes[i]!.attributes!.currentHP! = eri.value.eTeam.heroes[i]?.attributes!.maxHP! + pvAccumulatorEn[i]!
        }
    }

    private findHeroIdx(team: Team, hero: Hero): number {
        for (let i = 0; i < team.heroes.length; i++) {
            if (hero.id == team.heroes[i]?.id) {
                return i
            }
        }
        return -1
    }

    private getIdxOfTimeline(tl: ExpeditionStepTimestamp[], nowTime: number): number {
        for (let i = 0; i < tl.length; i++) {
            if (nowTime < Date.parse(tl[i]?.when!)) {
                return i
            }
        }
        return tl.length
    }

    async equip(e: Weapon | Armor | Omamori) {
        switch (this.inventoryType.value) {
            case EquipmentType.WeaponType:
                let w = e as Weapon
                let i1 = this.user.value!.inventory!.weapons.indexOf(w);
                if (i1 > -1) {
                    this.user.value!.inventory!.weapons.splice(i1, 1);
                }
                this.user.value!.inventory!.weapons.push(w)
                this.heroToEquip.value!.equipment!.weapon = w
                break;
            case EquipmentType.ArmorType:
                let a = e as Armor
                let i2 = this.user.value!.inventory!.armors.indexOf(a);
                if (i2 > -1) {
                    this.user.value!.inventory!.armors.splice(i2, 1);
                }
                this.user.value!.inventory!.armors.push(a)
                this.heroToEquip.value!.equipment!.armor = a
                break;
            case EquipmentType.OmamoriType:
                let o = e as Omamori
                let i3 = this.user.value!.inventory!.omamoris.indexOf(o);
                if (i3 > -1) {
                    this.user.value!.inventory!.omamoris.splice(i3, 1);
                }
                this.user.value!.inventory!.omamoris.push(o)
                this.heroToEquip.value!.equipment!.omamori = o
                break;
        }
        let params = [
            {
                id: "uid",
                value: `${this.user.value!.id}`,
            },
            {
                id: "hid",
                value: `${this.heroToEquip.value!.id}`,
            },
            {
                id: "etype",
                value: `${this.inventoryType.value}`,
            },
            {
                id: "eid",
                value: `${e.id}`,
            },
        ]
        this.heroToInspect.value = this.heroToEquip.value
        this.setHomeStatus(NavigationStatus.InspectHero)
        await fetchData<undefined>(ref(undefined), RequestType.Equip, params)
        await this.fetchInventory()
        await this.fetchTeam()
    }
}

export {
    NavigationHandler,
    NavigationStatus,
}
