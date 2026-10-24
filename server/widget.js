const LOCALSTORAGE_PREFIX = "nawst"
const NOW = new Date();

const GetLastVisited = function () {
    let _storedDate = window.localStorage.getItem(`${LOCALSTORAGE_PREFIX}_LastVisited`)
    return _storedDate == null ? new Date(_storedDate) : null
}

const DateWithinLastWeek = function (dt) {
    if (dt == null) {
        return false;
    }
    const dayOfWeek = NOW.getDay();
    const distanceToMonday = (dayOfWeek === 0) ? 6 : dayOfWeek - 1;
    const lastMonday = new Date(NOW);
    lastMonday.setDate(NOW.getDate() - distanceToMonday);
    lastMonday.setHours(0, 0, 0, 0);

    return dt.getTime() > lastMonday.getTime();
}

const nawst = function () {
    console.log("[NAWST] Loading... Please stand by.")
    const timezone = Intl.DateTimeFormat().resolvedOptions().timeZone;
    const domain = window.location.hostname;
    const path = window.location.pathname;

    const lastVisitedDateTime = GetLastVisited()
    const visitedThisWeek = DateWithinLastWeek(lastVisitedDateTime)

    window.localStorage.setItem(`${LOCALSTORAGE_PREFIX}_LastVisited`, NOW.toISOString());

    const scriptIngestURL = "{{ .scriptIngestURL }}";


    fetch(scriptIngestURL, {
        headers: {
            "Content-Type": "application/json",
        },
        method: "POST",
        body: JSON.stringify({
            tz: timezone, //timezone
            dm: domain, //domain, e.g. niceygy.net
            pt: path, //path, e.g. /home/article
            lw: visitedThisWeek, //visited since monday midnight
            vb: lastVisitedDateTime != null //visited at all
        })
    }).then((_) => console.log("[NAWST]: Sent\n[NAWST]: For more info, see https://niceygy.net/projects/nawst"))
    .catch((err) => console.log(`[NAWST] Response ERROR! '${err}'`))
}

if (document.readyState === "complete" || document.readyState === "interactive") {
    nawst()
} else {
    window.addEventListener("load", nawst);
}