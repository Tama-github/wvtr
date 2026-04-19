<script setup lang="ts">
    import { ref, onMounted, inject, watch } from 'vue'
    import { fetchData, getUserIDFromCookiesOrURLParams, global, RequestType } from "./model/utils.ts"
    import Header from "./components/Header.vue"
    import Body from "./components/Body.vue"
    import type { User } from "./model/types.ts"
    import type { VueCookies } from 'vue-cookies'

    const $cookies = inject<VueCookies>('$cookies');
    const userid = ref<string|null|undefined>(undefined)
    const user = ref<User|undefined>(undefined)
    const authUrl = ref<string|undefined>(undefined)
    const client_id = ref<string>("japan7")
    
    onMounted(async () => {
        userid.value = getUserIDFromCookiesOrURLParams($cookies)
        if (userid.value == undefined || userid.value == null) {
            const authServer = "https://auth.japan7.bde.enseeiht.fr";

            const resp = await fetch(`${authServer}/.well-known/openid-configuration`);
            const config = await resp.json();
            console.log(config);
            const params = new URLSearchParams();
            params.set("response_type", "code");
            params.set("client_id", client_id.value);
            params.set("redirect_uri", `${global.DOMAIN_NAME}/api/oidc/callback`);
            params.set("scope", "openid profile discord_id");
            authUrl.value = `${config.authorization_endpoint}?${params.toString()}`;
            window.location.replace(authUrl.value);
        } else {
            $cookies?.set("wvtrusrid", userid.value, '30d',undefined, undefined, true, "Strict")
            let urlParams = new URLSearchParams(window.location.search);
            if (urlParams.has('wvtrusrid')) {
                window.location.replace(global.DOMAIN_NAME);
            }
        }
    })

    watch(userid, async (nUser) => {
        if (nUser) {
            await fetchData<User>(user, RequestType.User, [{id: "id", value: `${userid.value}`}])
        }
    })
</script>

<template>
<div v-if="!user" class="page">
    <!-- <a v-if="authUrl" :href="authUrl">Login with OIDC</a> -->
    <p>loading auth...</p>
</div>
<div v-else class="page">
    <Header :user="user"/>
    <Body  :user="user"/>
</div>
</template>

