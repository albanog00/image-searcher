import { sequence } from "astro/middleware";
import { log } from "./log";

export const onRequest = sequence(log);