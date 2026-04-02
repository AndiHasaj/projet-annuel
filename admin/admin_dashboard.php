<?php
// Appel API pour les statistiques globales
$data = file_get_contents("http://localhost:8080/api/admin/stats");
$stats = json_decode($data, true);

// Fallback si le Go est éteint
if (!$stats) {
    $stats = ["revenus_24h" => 0, "nouveaux_pros" => 0, "signalements_box" => 0, "formations_attente" => 0, "total_mensuel" => 0, "croissance" => 0];
}
?>
<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <title>UCU Admin | Dashboard</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600&family=Montserrat:wght@700&display=swap" rel="stylesheet">
    <script>
        tailwind.config = { theme: { extend: { colors: { 'ucu-foret': '#1E4E4E', 'ucu-lagon': '#6A9A9A', 'ucu-grey': '#F8FAFC' }, fontFamily: { 'sans': ['Inter', 'sans-serif'], 'montserrat': ['Montserrat', 'sans-serif'] }, borderRadius: { 'ucu': '8px' } } } }
    </script>
</head>
<body class="bg-ucu-grey font-sans text-gray-800">
    <?php include '../includes/sidebar.php'; ?>
    <main class="lg:ml-72 p-12">
        <header class="flex justify-between items-center mb-12">
            <h2 class="font-montserrat font-bold text-3xl text-ucu-foret uppercase tracking-tighter">Console <span class="text-ucu-lagon italic">Générale</span></h2>
        </header>

        <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-12">
            <div class="bg-white p-6 rounded-ucu border border-gray-100 shadow-sm">
                <p class="text-[9px] font-bold text-gray-400 uppercase">Revenus (24h)</p>
                <p class="text-xl font-montserrat font-bold text-ucu-foret">+ <?= number_format($stats['revenus_24h'], 0, ',', ' ') ?> €</p>
            </div>
            <div class="bg-white p-6 rounded-ucu border border-gray-100 shadow-sm">
                <p class="text-[9px] font-bold text-gray-400 uppercase">Nouveaux Pros</p>
                <p class="text-xl font-montserrat font-bold text-ucu-lagon"><?= $stats['nouveaux_pros'] ?></p>
            </div>
            <div class="bg-white p-6 rounded-ucu border border-gray-100 shadow-sm">
                <p class="text-[9px] font-bold text-gray-400 uppercase">Signalements Box</p>
                <p class="text-xl font-montserrat font-bold text-red-500"><?= $stats['signalements_box'] ?></p>
            </div>
            <div class="bg-white p-6 rounded-ucu border border-gray-100 shadow-sm">
                <p class="text-[9px] font-bold text-gray-400 uppercase">Formations en attente</p>
                <p class="text-xl font-montserrat font-bold text-ucu-foret"><?= $stats['formations_attente'] ?></p>
            </div>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-10">
            <section class="bg-ucu-foret p-8 rounded-ucu text-white shadow-xl">
                <h3 class="font-montserrat font-bold text-[10px] uppercase text-ucu-lagon mb-6 tracking-widest">Performances Financières</h3>
                <div class="flex items-end gap-4 mb-6">
                    <p class="text-4xl font-montserrat font-bold"><?= number_format($stats['total_mensuel'], 0, ',', ' ') ?> €</p>
                    <p class="text-[10px] text-green-400 font-bold mb-2 uppercase">▲ <?= $stats['croissance'] ?>% ce mois</p>
                </div>
                <button onclick="location.href='admin_finance.php'" class="w-full mt-8 py-3 bg-ucu-lagon rounded font-bold text-[9px] uppercase tracking-widest hover:bg-white hover:text-ucu-foret transition-all">Analyse détaillée</button>
            </section>
        </div>
    </main>
</body>
</html>