<script setup lang="ts">
    import { inject } from "vue";
import type { Hero } from "../model/types.ts"
    import { global, NavigationHandler } from "../model/utils.ts"

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const hero = navigationHandler.getHeroToInspect()
    console.log(hero.value)
</script>

<template>
    <div v-if="hero" class="inspect-hero">
        <h1>{{ hero.name }}<label style="font-size: medium;"> lvl.{{ hero.attributes.level }}</label></h1>
        <div class="row">
            <div>
                <img :src="hero.imageUrl" width="200"/>
            </div>
            <div class="column">
                <div>
                    <div>Class: {{ hero.heroClass.name }}</div>
                </div>
                <div>
                    <div class="column" style="display: inline-flex;">
                    <label>hp : {{ hero.attributes.currentHP }}/{{ hero.attributes.maxHP }} </label>
                    <progress :max="hero.attributes.maxHP" :value="hero.attributes.currentHP" class="hero-progress"/>
                    <label>xp : {{ hero.attributes.currentXP }}/{{ hero.attributes.xpBeforLvlUp }} </label>
                    <progress :max="hero.attributes.xpBeforLvlUp" :value="hero.attributes.currentXP" class="hero-progress"/>
                </div>
                </div>
                <div>
                    <div>Strength ({{ hero.attributes.sgt.toFixed(2) }}) : {{ hero.attributes.strength }}</div>
                    <div>Intelligence ({{ hero.attributes.igt.toFixed(2) }}) : {{ hero.attributes.intelligence }}</div>
                    <div>Dexterity ({{ hero.attributes.dgt.toFixed(2) }}) : {{ hero.attributes.dexterity }}</div>
                    <div>Luck ({{ hero.attributes.lgt.toFixed(2) }}) : {{ hero.attributes.luck }}</div>
                </div>
                <div>
                    <div>Resitances :</div>
                    <div>Blunt : {{ hero.attributes.blunt }}</div>
                    <div>Pierce : {{ hero.attributes.pierce }}</div>
                    <div>Slash : {{ hero.attributes.slash }}</div>
                    <div>Fire : {{ hero.attributes.fire }}</div>
                    <div>Frost : {{ hero.attributes.frost }}</div>
                    <div>Lightning : {{ hero.attributes.lighting }}</div>
                </div>
            </div>
            <div class="column">
                <label>Skills</label>
                <div class="raw">
                    <div v-if="hero.uniqueSkill" class="column">
                        <label>{{ hero.uniqueSkill.name }}</label>
                        <img v-if="hero.uniqueSkill.image_url !== ''" :src="global.DOMAIN_NAME + hero.uniqueSkill.image_url"/>
                    </div>
                    <div v-if="hero.activeSkill" class="column">
                        <label>{{ hero.activeSkill.name }}</label>
                        <img v-if="hero.activeSkill.image_url !== ''" :src="global.DOMAIN_NAME + hero.activeSkill.image_url"/>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div v-else>
        no hero to inspect
    </div>
</template>