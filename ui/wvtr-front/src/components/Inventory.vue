<script setup lang="ts">
    import { inject, ref } from "vue";
import { EquipmentType, type Armor, type Equipable, type Hero, type Omamori, type Weapon } from "../tools/types.ts"
    import { fetchData, global, RequestType } from "../tools/utils.ts"
import type { NavigationHandler } from "@/tools/navigationHandler.ts";
import InspectEquipment from "./InspectEquipment.vue";

    const navigationHandler = inject<NavigationHandler>('navigationHandler')!
    const user = navigationHandler.getUser()
    const inventoryType = navigationHandler.getInventoryType()
    const equipmentToInspect = ref<Weapon | Armor | Omamori | undefined>(undefined)
    const errorMsg = ref("")

    function clickOnArmor(e: Armor) {
        equipmentToInspect.value = e
    }
    
    function clickOnWeapon(e: Weapon) {
        equipmentToInspect.value = e
    }

    function clickOnOmamori(e: Omamori) {
        equipmentToInspect.value = e
    }

    function onclick() {
        navigationHandler.equip(equipmentToInspect.value!)
    }
    
</script>

<template>
    <div class="column">
        <h1>Select an equipment</h1>
        <div style="display: flex; align-items: center; justify-content: center;">
            <div>
                <button v-on:click="onclick()">Equip</button>
                <p>{{ errorMsg }}</p>
            </div>
        </div>
        <div class="row" v-if="inventoryType == EquipmentType.WeaponType">
            <div class="row">
                <div v-on:click="clickOnWeapon(e)" v-for="e in user?.inventory.weapons">
                    <img :src="e.iconURL" width="50px">
                </div>
            </div>
            <InspectEquipment :weapon="(equipmentToInspect as Weapon)" :armor="undefined" :omamori="undefined" />
        </div>
        <div class="row" v-else-if="inventoryType == EquipmentType.ArmorType">
            <div class="row">
                <div v-on:click="clickOnArmor(e)" v-for="e in user?.inventory.armors">
                    <img :src="e.iconURL" width="50px">
                </div>
            </div>
            <InspectEquipment :weapon="undefined" :armor="(equipmentToInspect as Armor)" :omamori="undefined"/>
        </div>
        <div class="row" v-else-if="inventoryType == EquipmentType.OmamoriType">
            <div class="row">
                <div v-on:click="clickOnOmamori(e)" v-for="e in user?.inventory.omamoris">
                    <img :src="e.iconURL" width="50px">
                </div>
            </div>
            <InspectEquipment :weapon="undefined" :armor="undefined" :omamori="(equipmentToInspect as Omamori)"/>
        </div>
    </div>
</template>