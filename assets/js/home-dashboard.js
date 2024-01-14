const app_url = $("#app_url").val();
const chartBarVal = $("#chart-bar-val").val().split("_");
const chartDoughnutVal = $("#chart-doughnut-val").val().split("_");

const barChartDashboard = (type, elname, datasets = []) => {
    const ctx = document.getElementById(elname);
    new Chart(ctx, {
        type: type,
        data: {
            labels: ['Create', 'Delete'],
            datasets: [{
                label: '# of Votes',
                data: [datasets[0], datasets[1]],
                borderWidth: 5
            }]
        },
        options: {
            scales: {
                y: {
                    beginAtZero: true
                }
            }
        }
    });
}

$(document).ready(function() {
    barChartDashboard("bar", "bar-chart", chartBarVal);
    barChartDashboard("doughnut", "doughnut-chart", chartDoughnutVal);
});