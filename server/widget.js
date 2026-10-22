const LOCALSTORAGE_PREFIX = "nawst"
const NOW = new Date();

const GetLastVisited = function () {
    let _storedDate = window.localStorage.getItem(`${LOCALSTORAGE_PREFIX}_LastVisited`)
    if (_storedDate) {
        return new Date(_storedDate);
    } else {
        return null
    }
}

const SetLastVisited = function (dt) {
    window.localStorage.setItem(`${LOCALSTORAGE_PREFIX}_LastVisited`, dt.toISOString());
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

window.onload = function () {
    const timezone = Intl.DateTimeFormat().resolvedOptions().timeZone;
    const domain = window.location.hostname;
    const path = window.location.pathname;

    const lastVisitedDateTime = GetLastVisited()
    const visitedThisWeek = DateWithinLastWeek(lastVisitedDateTime)

    console.log(`[nawst]: tz=${timezone}`)
    console.log(`[nawst]: domain=${domain}`)
    console.log(`[nawst]: path=${path}`)
    console.log(`[nawst]: lastVisit=${lastVisitedDateTime}`)
    console.log(`[nawst]: visitedThisWeek=${visitedThisWeek}`)

    SetLastVisited(NOW)

    const scriptIngestURL = "{{ .scriptIngestURL }}";


    fetch(scriptIngestURL, {
        method: "POST",
        body: JSON.stringify({
            tz: timezone, //timezone
            dm: domain, //domain, e.g. niceygy.net
            pt: path, //path, e.g. /home/article
            lw: visitedThisWeek, //visited since monday midnight
            vb: lastVisitedDateTime != null //visited at all
        })
    }).then((_) => console.log("[NAWST]: Sent\n[NAWST]: For more info, see https://niceygy.net/projects/nawst"))
}