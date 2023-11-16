import type { ApiResponse } from "./responses";

export async function api<T>(input: RequestInfo, init?: RequestInit | undefined): Promise<ApiResponse<T>> {
    performance.mark("apiStart")
    const res = await fetch(`/api${input}`, init);
    performance.mark("apiEnd")

    const elapsed = performance.measure("apiDuration", {
        start: "apiStart",
        end: "apiEnd",
        duration: undefined,
    });

    console.log(init?.method, res.url, `${elapsed.duration.toFixed(3)}ms`);

    return res;
}