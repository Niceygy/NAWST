window.onload = function() {
    const timezone = Intl.DateTimeFormat().resolvedOptions().timeZone;
    const domain = window.location.hostname;
    const path = window.location.pathname;
    const visitedBefore = window.localStorage.getItem("niwst_visitedSiteBefore") == "true";

    const scriptIngestURL = "{{ .scriptIngestURL }}";

    window.localStorage.setItem("niwst_visitedSiteBefore", "true");

    fetch(scriptIngestURL, {
        method: "POST",
        body: JSON.stringify({
            tz: timezone,
            dm: domain,
            pt: path,
            vb: visitedBefore
        })
    }).then((_) => console.log("[NAWST]: Sent\n[NAWST]: For more info, see https://niceygy.net/projects/nawst"))
}