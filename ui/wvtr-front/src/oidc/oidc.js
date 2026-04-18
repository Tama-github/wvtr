import express from "express";

const authServer = "https://auth.japan7.bde.enseeiht.fr";
const clientId = /** @type {string} */ (process.env.CLIENT_ID);
const clientSecret = /** @type {string} */ (process.env.CLIENT_SECRET);

const resp = await fetch(`${authServer}/.well-known/openid-configuration`);
const config = await resp.json();
console.log(config);

const app = express();

app.get("/", (req, res) => {
    const params = new URLSearchParams();
    params.set("response_type", "code");
    params.set("client_id", clientId);
    params.set("redirect_uri", "http://localhost:3000/api/oidc/callback");
    params.set("scope", "openid profile discord_id");
    const authUrl = `${config.authorization_endpoint}?${params.toString()}`;
    res.send(`<a href="${authUrl}">Login with OIDC</a>`);
});

app.get("/api/oidc/callback", async (req, res) => {
    const code = /** @type {string} */ (req.query.code);

    const params = new URLSearchParams();
    params.set("grant_type", "authorization_code");
    params.set("code", code);
    params.set("client_id", clientId);
    params.set("client_secret", clientSecret);
    params.set("redirect_uri", "http://localhost:3000/api/oidc/callback");
    const tokenResp = await fetch(config.token_endpoint, {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded",
        },
        body: params.toString(),
    });
    const tokenData = await tokenResp.json();

    const userInfoResp = await fetch(config.userinfo_endpoint, {
        headers: {
            Authorization: `Bearer ${tokenData.access_token}`,
        },
    });
    const userInfo = await userInfoResp.json();

    res.send({ tokenData, userInfo });
});

app.listen(3000, () => {
    console.log("Go to http://localhost:3000 to start");
});