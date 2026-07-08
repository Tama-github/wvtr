<script setup lang="ts">
    import { inject } from "vue";
import { EquipmentType, type Armor, type Hero, type Omamori, type Weapon } from "../tools/types.ts"
    import { global } from "../tools/utils.ts"
import type { NavigationHandler } from "@/tools/navigationHandler.ts";
import Hoverable from "./Hoverable.vue";
import InspectEquipment from "./InspectEquipment.vue";
import WeaponHoverable from "./HoverComponents/WeaponHoverable.vue";
import SkillHoverable from "./HoverComponents/SkillHoverable.vue";
import OmamoriHoverable from "./HoverComponents/OmamoriHoverable.vue";
import ArmorHoverable from "./HoverComponents/ArmorHoverable.vue";

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const hero = navigationHandler.getHeroToInspect()
    const user = navigationHandler.getUser()
    console.log(hero.value)

    function clickEquiment(type: EquipmentType) {
        navigationHandler.setInventoryVue(type, hero.value!)
    }

</script>

<template>
    <div v-if="hero" class="inspect-hero">
        <h1>{{ hero.name }}<label style="font-size: medium;"> lvl.{{ hero.attributes!.level }}</label></h1>
        <div class="row">
            <div>
                <img :src="hero.imageUrl" width="200"/>
            </div>
            <div class="column">
                <div>
                    <div>Class: {{ hero.heroClass!.name }}</div>
                </div>
                <div>
                    <div class="column" style="display: inline-flex;">
                    <label>hp : {{ hero.attributes!.currentHP.toFixed(0) }}/{{ hero.attributes!.maxHP.toFixed(0) }} </label>
                    <progress :max="hero.attributes!.maxHP.toFixed(0)" :value="hero.attributes!.currentHP.toFixed(0)" class="hero-progress"/>
                    <label>xp : {{ hero.attributes!.currentXP.toFixed(0) }}/{{ hero.attributes!.xpBeforLvlUp.toFixed(0) }} </label>
                    <progress :max="hero.attributes!.xpBeforLvlUp.toFixed(0)" :value="hero.attributes!.currentXP.toFixed(0)" class="hero-progress"/>
                </div>
                </div>
                <div>
                    <label>Attributes:</label>
                    <div>Strength ({{ hero.attributes!.sgt.toFixed(2) }}) : {{ hero.attributes!.strength }}</div>
                    <div>Intelligence ({{ hero.attributes!.igt.toFixed(2) }}) : {{ hero.attributes!.intelligence }}</div>
                    <div>Dexterity ({{ hero.attributes!.dgt.toFixed(2) }}) : {{ hero.attributes!.dexterity }}</div>
                    <div>Luck ({{ hero.attributes!.lgt.toFixed(2) }}) : {{ hero.attributes!.luck }}</div>
                </div>
                <div>
                    <div>Resitances :</div>
                    <div class="row" style="gap: 10px;">
                        <div>
                            <div>Blunt:</div>
                            <div>Pierce:</div>
                            <div>Slash:</div>
                        </div>
                        <div>
                            <div>{{ hero.attributes!.blunt }}</div>
                            <div>{{ hero.attributes!.pierce }}</div>
                            <div>{{ hero.attributes!.slash }}</div>
                        </div>
                        <div>
                            <div>Fire:</div>
                            <div>Frost:</div>
                            <div>Lightning:</div>
                        </div>
                        <div>
                            <div>{{ hero.attributes!.fire }}</div>
                            <div>{{ hero.attributes!.frost }}</div>
                            <div>{{ hero.attributes!.lighting }}</div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="column">
                <label>Skills</label>
                <div class="raw">
                    <div class="column">
                        <div> Unique skill </div>
                        <SkillHoverable v-if="hero.uniqueSkill" :skill="hero.uniqueSkill">
                            <div class="column">
                                <img v-if="hero.uniqueSkill.image_url !== ''" :src="hero.uniqueSkill.image_url" width="75"/>
                            </div>
                        </SkillHoverable>
                    </div>
                    <div class="column">
                        <div> Active skill </div>
                        <SkillHoverable v-if="hero.activeSkill" :skill="hero.activeSkill">
                            <div class="column">
                                <img v-if="hero.activeSkill.image_url !== ''" :src="hero.activeSkill.image_url"/>
                            </div>
                        </SkillHoverable>
                        <div v-else class="column">
                            <img width="75" :src="global.NO_EQUIPMENT"/>
                        </div>
                    </div>
                </div>
            </div>
            <div class="column">
                <label>Equipment</label>
                <div class="raw">
                    <div class="column">
                        <label>Weapon</label>
                        <WeaponHoverable v-if="hero.equipment!.weapon" :weapon="hero.equipment!.weapon">
                            <img v-on:click="clickEquiment(EquipmentType.WeaponType)" width="75" :src="hero.equipment!.weapon.iconURL"/>
                        </WeaponHoverable>
                        <img v-else v-on:click="clickEquiment(EquipmentType.WeaponType)" width="75" :src="global.NO_EQUIPMENT"/>
                    </div>

                    <div class="column">
                        <label>Armor</label>
                        <ArmorHoverable v-if="hero.equipment!.armor" :armor="hero.equipment!.armor">
                            <img v-on:click="clickEquiment(EquipmentType.ArmorType)" width="75" :src="hero.equipment!.armor?.iconURL"/>
                        </ArmorHoverable>
                        <img v-else v-on:click="clickEquiment(EquipmentType.ArmorType)" width="75" :src="global.NO_EQUIPMENT"/>
                    </div>
                    <div class="column">
                        <label>Omamori</label>
                        <OmamoriHoverable v-if="hero.equipment!.omamori" :omamori="hero.equipment!.omamori">
                            <img v-on:click="clickEquiment(EquipmentType.OmamoriType)" width="75" :src="hero.equipment!.omamori?.iconURL"/>
                        </OmamoriHoverable>
                        <img v-else v-on:click="clickEquiment(EquipmentType.OmamoriType)" width="75" :src="global.NO_EQUIPMENT"/>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div v-else>
        no hero to inspect
    </div>
</template>