
type Hero = {
    id: number;
    name: string;
    imageUrl: string;
    level: number;
    currentXP: number;
    xpBeforLvlUp: number;
    currentHP: number;
    maxHP: number;
};

type Team = {
    id?: number;
    heroes?: Hero[];
};

enum EncounterState {
    Home = 1,
    Travel,
    Fight,
    Neutral,
    Error,
}

type ExpeditionStepResolveInfo = {
    stepInfos: string,
    stepEndAt: string,
    stepState: EncounterState,
}

type GameState = {
    state: EncounterState,
    wTeam: Team | null,
    eTeam: Team | null,
}

type User = {
    id: number,
    name: string,
    state: GameState,
    currentTeam: Team,
    lastActionTime: string,
    ownedHeroes: Hero[],
}

type CurrentStepRequestMessage = {
    id: number
    time: number
}

type Waifu = {
    id: string,
    id_al: string,
    name_user_preferred: string,
    image_large: string,
    rank: string,
}

export type {
    Hero,
    Team,
    GameState,
    User,
    ExpeditionStepResolveInfo,
    CurrentStepRequestMessage,
    Waifu,
};

export {
    EncounterState,
}

