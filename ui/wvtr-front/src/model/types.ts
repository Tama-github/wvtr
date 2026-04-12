
type Waifu = {
    id?: number;
    imageUrl?: string;
};

type Team = {
    id?: number;
    waifus?: Waifu[];
};

enum EncounterState {
    Home = 1,
    Travel,
    Fight,
    Neutral,
    Error,
}

type GameState = {
    isBusy: boolean;
    state: EncounterState;
    wTeam: Team | null;
    eTeam: Team | null;
}

type User = {
    id: number,
    name: string,
    state: GameState,
    currentTeam: Team,
    lastActionTime: Date,
}

export type {
    Waifu,
    Team,
    GameState,
    User,
};

export {
    EncounterState,
}

