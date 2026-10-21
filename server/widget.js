window.onload = function() {
    const timezone = Intl.DateTimeFormat().resolvedOptions().timeZone;
    const domain = window.location.hostname;
    const path = window.location.pathname;
    const visitedBefore = window.localStorage.getItem("nawst_visitedSiteBefore") == "true";
    let visitorFingerprint = window.localStorage.getItem("nawst_userPrint")

    if (visitorFingerprint == null) {
        visitorFingerprint = Date.now().toString(36) + Math.random().toString(36).substr(2);
    }

    const scriptIngestURL = "{{ .scriptIngestURL }}";

    window.localStorage.setItem("niwst_visitedSiteBefore", "true");

    fetch(scriptIngestURL, {
        method: "POST",
        body: JSON.stringify({
            tz: timezone,
            dm: domain,
            pt: path,
            vb: visitedBefore,
            fp: visitorFingerprint,
        })
    }).then((_) => console.log("[NAWST]: Sent\n[NAWST]: For more info, see https://niceygy.net/projects/nawst"))
}