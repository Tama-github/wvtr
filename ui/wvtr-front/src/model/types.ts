type HeroAttributes = {

    level: number
    currentXP: number
    xpBeforLvlUp: number
    currentHP: number

    //Attributes
    maxHP: number
    strength: number
    intelligence: number
    dexterity: number
    luck: number

    //Growth rate
    hpgt: number
    sgt: number
    igt: number
    dgt: number
    lgt: number

    // Resistances
    blunt: number
    pierce: number
    slash: number
    fire: number
    frost: number
    lighting: number
}

type HeroClass = {
    name: string
    descritpion: string
}

enum SkillType {
    Unique = 0,
    Active,
}

enum SkillID {
    Lucky = 0,
    GoodRest,
    SecondWind,
    Prodigy,
    Berserk,
    Trickster,
    FastLearner,
    ElementalCursed,
    PhysicalCursed,
}

type Skill = {
    identifier: SkillID
    name: string
    skill_type: SkillType
    image_url: string
    description: string
    weight: number
}

type Hero = {
    id: number;
    imageUrl: string
    name: string
    heroClass: HeroClass
    rank: string
    attributes: HeroAttributes
    uniqueSkill: Skill
    activeSkill: Skill

    // info that we save to request nanapi if we need to.
    id_w: string
    id_al: number
};

type Team = {
    id: number;
    heroes: Hero[];
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
    id: number
    name: string
    state: GameState
    currentTeam: Team
    lastActionTime: string
    ownedHeroes: Hero[]
    discord_id: string
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

