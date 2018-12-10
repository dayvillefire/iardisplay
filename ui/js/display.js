// IAR Display logic

var incidentCache = {};

$(document).ready(function () {
    // Initial load
    populateIncidents();
    populateOnDuty();
    populateResponding();
    populateCad();

    // Set repeating
    setInterval(populateResponding, 10500);
    setInterval(populateOnDuty, 10250)
    setInterval(populateIncidents, 10000);
    setInterval(populateCad, 10750);
});

function populateCad() {
    $.getJSON("/api/cad/current", {}, function (response) {
        //$.getJSON("/api/cad/cleared/12-8-2018", {}, function (response) {
        var body = '';
        console.log("populateCad(): " + JSON.stringify(response));
        for (var i in response) {
            // Iterate through each call
            //console.log(JSON.stringify(response[i]));
            body += '<tr>';
            body +=
                '<td>' + i + '</td>' +
                '<td>' + response[i]['priority'] + '</td>' +
                '<td>' + response[i]['location'] + '</td>' +
                '<td>' + response[i]['call_type'] + '</td>';
            // Figure out unit status
            var unitsresponding = [];
            var unitsonscene = [];
            var unitscleared = [];
            for (var unit in response[i].units) {
                // TODO: FIXME: XXX: refactor this logic into something configured on the server side
                if (unit.startsWith('STA') || unit.startsWith('RES') || unit.endsWith('OFF') || unit.startsWith('KB')) {
                    continue;
                }
                if (response[i].units[unit]['status'] == 'RESPONDING') {
                    unitsresponding.push(highlightOurUnits(unit));
                }
                if (response[i].units[unit]['status'] == 'ON SCENE') {
                    unitsonscene.push(highlightOurUnits(unit));
                }
                if (response[i].units[unit]['status'] == 'Available') {
                    unitscleared.push(highlightOurUnits(unit));
                }
            }
            // Make non-filled fields say 'none'
            if (unitsresponding.length == 0) {
                unitsresponding.push("none");
            }
            if (unitsonscene.length == 0) {
                unitsonscene.push("none");
            }
            if (unitscleared.length == 0) {
                unitscleared.push("none");
            }
            body += '<td>' + unitsresponding.join(',') + '</td>' +
                '<td>' + unitsonscene.join(',') + '</td>' +
                '<td>' + unitscleared.join(',') + '</td>' +
                '</tr>';
        }
        if (body == '') {
            body = '<tr><td colspan="7"><div align="center">No active cases</div></td></tr>';
        }
        $('#cad').html(body);
        //console.log(body);
    });
}

function populateIncidents() {
    $.getJSON("/api/iar/incidents", {}, function (response) {
        //var body = '';
        $('#incidents').html('');
        x = response.incidents.length;
        if (x > 5) {
            x = 5;
        }
        for (var i = 0; i < x; i++) {
            var id = response.incidents[i].Id;
            var body = '<tr><td>';
            if (i == 0) {
                body += "<b>";
            }
            body += response.detail[id + 0].ArrivedOn.replace('T', ' ');
            if (i == 0) {
                body += "</b>";
            }
            body += "</td><td>";
            if (i == 0) {
                body += "<b>";
            }
            body += response.detail[id + 0].Body;
            if (i == 0) {
                body += "</b>";
            }
            body += '</td></tr>';
            $('#incidents').append(body);
        }
    });
    //body += '<tr><td>' + response[i].Body + '</td></tr>';
}

function populateOnDuty() {
    $.getJSON("/api/iar/schedule", {}, function (response) {
        var body = '';
        x = response.length;
        for (i = 0; i < x; i++) {
            body += '<tr>' +
                '<td>' + response[i].MemberName + '</td>' +
                '<td>' + response[i].MemberCat + '</td>' +
                '<td>' + response[i].InStationOrHome + '</td>' +
                '<td>' + response[i].MemberStation + '</td>' +
                '<td>' + response[i].UntilAt + '</td>' +
                '</tr>';
        }
        if (body == '') {
            body = '<tr><td colspan="5"><div align="center">No one scheduled</div></td></tr>';
        }
        $('#schedule').html(body);
        //console.log(body);
    });
}

function populateResponding() {
    $.getJSON("/api/iar/responding", {}, function (response) {
        var body = '';
        //console.log("populateResponding(): " + JSON.stringify(response));
        x = response.length;
        for (i = 0; i < x; i++) {
            console.log(JSON.stringify(response[i]));
            body += '<tr>' +
                '<td>' + response[i].MemberFName + '</td>' +
                '<td>' + response[i].MemberCat + '</td>' +
                '<td>' + response[i].MemberStation + '</td>' +
                '<td>' + response[i].CallingTime + '</td>' +
                '<td>' + response[i].ETA + '</td>' +
                '</tr>';
        }
        if (body == '') {
            body = '<tr><td colspan="5"><div align="center">No one responding</div></td></tr>';
        }
        $('#responding').html(body);
        //console.log(body);
    });
}

function highlightOurUnits(unit) {
    // TODO: FIXME: should be configured server side rather than hardcoded
    if (!unit.endsWith('63')) {
        return unit
    }
    return '<b><font color="#f00">' + unit + '</font></b>';
}