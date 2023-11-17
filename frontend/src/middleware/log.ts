import { defineMiddleware } from "astro/middleware";

export const log = defineMiddleware((context, next) => {
    performance.mark("requestStart")
    next();
    performance.mark("requestEnd")

    const elapsed = performance.measure("requestDuration", {
        start: "requestStart",
        end: "requestEnd",
        duration: undefined,
    });

    console.log(context.request.method, context.request.url, `served in ${elapsed.duration.toFixed(3)}ms`);
});