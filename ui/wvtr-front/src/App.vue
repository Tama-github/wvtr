<script setup lang="ts">
    import { ref, onMounted, inject, watch } from 'vue'
    import { fetchData, getUserIdFromCookies, getUserIdFromRequestParams, global, isUserIdInCookies, isUserIdInRequestParams, RequestType } from "./model/utils.ts"
    import Header from "./components/Header.vue"
    import Body from "./components/Body.vue"
    import type { User } from "./model/types.ts"
    import type { VueCookies } from 'vue-cookies'

    const $cookies = inject<VueCookies>('$cookies');
    const user = ref<User|undefined>(undefined)
    const authUrl = ref<string|undefined>(undefined)
    
    onMounted(async () => {
        if ($cookies && isUserIdInRequestParams()) { // we got here only after authantification has been done
            let uidstring = getUserIdFromRequestParams()
            $cookies.set("wvtrusrid", uidstring, '30d',undefined, undefined, true, "Strict")
            window.location.replace(global.DOMAIN_NAME);
        } else if ($cookies && isUserIdInCookies($cookies)) { // client has already got here but we need to check if it matches the database
            await fetchData<User>(user, RequestType.User, [{id: "id", value: `${getUserIdFromCookies($cookies)}`}])
        } else { // we don't know the user and need to ask auth for it
            await requestToAuth()
        }
    })

    async function requestToAuth() {
        const authServer = "https://auth.japan7.bde.enseeiht.fr";
        const client_id = ref<string>("japan7")

        const resp = await fetch(`${authServer}/.well-known/openid-configuration`);
        const config = await resp.json();
        console.log("uidfcookie = " + getUserIdFromCookies($cookies!))
        console.log(config);
        const params = new URLSearchParams();
        params.set("response_type", "code");
        params.set("client_id", client_id.value);
        params.set("redirect_uri", `${global.DOMAIN_NAME}/api/oidc/callback`);
        params.set("scope", "openid profile discord_id");
        authUrl.value = `${config.authorization_endpoint}?${params.toString()}`;
        window.location.replace(authUrl.value);
    }

    watch(user, async (nUser) => {
        console.log("user.id = " + nUser?.id)
        // we got here only if cookies have a wvtrusrid and we fetched a user
        let uidfcookie = getUserIdFromCookies($cookies!)
        if (nUser && uidfcookie == nUser.id+"") { // check integrity of the cookie and the user
            // all good
        } else if (!nUser || (nUser && uidfcookie != nUser.id+"")) {
            // client think it has the right uid but it is wrong auth again
            requestToAuth()
        } else {
            console.log("problem")
            console.log(user)
            console.log(uidfcookie)
        }
    })

</script>

<template>
<div v-if="!user" class="page">
    <!-- <a v-if="authUrl" :href="authUrl">Login with OIDC</a> -->
    <p>loading auth...</p>
    <!-- <p v-else>loading auth...</p> -->
</div>
<div v-if="user" class="page">
    <Header :user="user"/>
    <Body  :user="user"/>
</div>
</template>

