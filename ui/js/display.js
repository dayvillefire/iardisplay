// IAR Display logic

var config = {};
var incidentCache = {};

$(document).ready(function () {
    // Load local UI config before we go any further
    $.getJSON("/api/config", {}, function (response) {
        config = response;
        console.log(config);
    });

    setInterval(showClock, 1000);

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

function showClock() {
    var d = new Date();
    d.setMinutes(d.getMinutes() - d.getTimezoneOffset());
    $('#clock').html(d.toISOString().slice(0, 19).replace('T', ' '));
}

function populateCad() {
    //$.getJSON("/api/cad/cleared/12-10-2018", {}, function (response) {
    $.getJSON("/api/cad/current", {}, function (response) {
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
                if (!isUnitValid(unit)) {
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
            // Use timestamp, but remove annoying 'T' separator
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

function isUnitValid(unit) {
    for (var i = 0; i < config.ignorePatterns.length; i++) {
        var p = config.ignorePatterns[i];
        if (p == unit) {
            return false;
        }
        if (p.endsWith('*')) {
            if (unit.startsWith(p.replace('*', ''))) {
                return false;
            }
        }
        if (p.startsWith('*')) {
            if (unit.endsWith(p.replace('*', ''))) {
                return false;
            }
        }
    }
    return true;
}

function highlightOurUnits(unit) {
    if (config.unitSuffix == '') {
        return unit;
    }
    if (!unit.endsWith(config.unitSuffix)) {
        return unit;
    }
    return '<b><font color="#f00">' + unit + '</font></b>';
}