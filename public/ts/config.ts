// @ts-ignore
export let env = process.env.NODE_ENV || "development";

export let baseURL = env == "development" ? "http://localhost:8080" : "http://localhost:8080";