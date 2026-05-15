
enum AffixType {
    Life = 0,
}

enum HeroTakeDamageStatus {
    TookDamage = 1 << 0,
    Dodged = 1 << 1,
    Blocked = 1 << 2,
    Died = 1 << 3,
    Crit = 1 << 4,
}

enum TargetType {
    Self = 1,
    Enemy,
    Friends,
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

enum EncounterState {
    Home = 1,
    Travel,
    Fight,
    Neutral,
    Report,
    Error,
}

enum CurrencyType {
    Gold = 0,
    MScrap,
    LSCrap,
    CScrap,
}

enum EquipmentType {
    WeaponType = 0,
    ArmorType,
    OmamoriType,
}

type Damage = {
    slashDmg: number,
    bluntDmg: number,
    pierceDmg: number,
    fireDmg: number,
    frostDmg: number,
    lightningDmg: number,
}

type StatsRange = {
    min: number
    max: number
    value: number
}

type Affix = {
    name: string
    ranges: StatsRange[]
    type: AffixType
}

type Storable = {
    name: string,
    iconURL: string,
}

type Currency = {
    type: CurrencyType,
    name: string,
    iconURL: string,
}

type CurrencyOwned = {
    numberOwned: number,
    currency: Currency,
}

type Equipable = Storable & {
    realWeightScore: number,
    affixes: Affix[],
}

type Weapon = Equipable & {
    id: number,
    baseDamage: Damage,
    baseCritRate: StatsRange,
    baseAttackSpeed: StatsRange,
}

type Armor = Equipable & {
    id: number,
    blockScore: StatsRange,
    evadeScore: StatsRange,
    baseResistancesRange: Damage,
}

type Omamori = Equipable & {
    id: number,
}

type HeroEquipment = {
    weapon: Weapon
    armor: Armor
    omamori: Omamori
}

type Inventory = {
    weapons: Weapon[]
    armors: Armor[]
    omamoris: Omamori[]
    currencies: CurrencyOwned[]
}

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

    //Defense
    blockScore: number,
    evadeScore: number,

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
    class_icon_url: string
}

type Reward = {
    xp: number,
    loot: Inventory,
}

type FieldActionDesc = {
    fromH: Hero
    usedSKill: Skill
    targetH: Hero
    targetStatus: HeroTakeDamageStatus
    fromPVChange: number
    targetPVChange: number
}

type ExpeditionStepTimestamp = {
    when: string, // time
    what: string,
    whatAction: FieldActionDesc,
}

type ExpeditionStepResolveInfo = {
    stepState: EncounterState,
    timeline: ExpeditionStepTimestamp[],
    eTeam: Team | null
}

type ExpeditionDB = {
    identifier: string,
    startedAt: string,
    whatHappened: ExpeditionStepResolveInfo[],
    ExpeditionRewards: Reward,
}

type GameState = {
    id: number
    state: EncounterState,
    currentExpedition: ExpeditionDB,
}

type User = {
    id: number
    name: string
    state: GameState
    inventory: Inventory
    currentTeam: Team
    lastActionTime: string // time
    ownedHeroes: Hero[]
    discord_id: string
}

type CurrentStepRequestMessage = {
    id: number
    time: number
}

type Skill = {
    identifier: SkillID
    name: string
    skill_type: SkillType
    target_type: TargetType
    recuperation_duration: number
    image_url: string
    description: string
}

type Hero = {
    id: number;
    imageUrl: string
    name: string
    heroClass: HeroClass
    rank: string
    attributes: HeroAttributes

    // skills
    weaponAttack: Skill
    uniqueSkill: Skill
    activeSkill: Skill

    // Items
    equipment: HeroEquipment

    // info that we save to request nanapi if we need to.
    id_w: string
    id_al: number
};

type Team = {
    id: number;
    heroes: Hero[];
};

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
    ExpeditionStepTimestamp,
    CurrentStepRequestMessage,
    ExpeditionDB,
    Waifu,
    FieldActionDesc,
    Inventory,
    Equipable,
    Weapon,
    Armor,
    Omamori,
};

export {
    EncounterState,
    HeroTakeDamageStatus,
    EquipmentType,
}

