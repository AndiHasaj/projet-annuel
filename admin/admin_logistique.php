<?php
// On récupère la liste des box via l'API
$response = file_get_contents("http://localhost:8080/api/admin/logistique");
$boxes = json_decode($response, true) ?? [];
?>
<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <title>UCU Admin | Logistique</title>
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
            <h2 class="font-montserrat font-bold text-3xl text-ucu-foret uppercase tracking-tighter">Parc de <span class="text-ucu-lagon italic">Conteneurs</span></h2>
        </header>

        <section class="bg-white rounded-ucu shadow-sm border border-gray-100 overflow-hidden">
            <table class="w-full text-left text-[11px]">
                <thead class="bg-ucu-grey border-b border-gray-100">
                    <tr class="text-gray-400 font-montserrat uppercase text-[9px] tracking-widest">
                        <th class="px-8 py-5">Matricule</th>
                        <th class="px-8 py-5">Localisation</th>
                        <th class="px-8 py-5">État de charge</th>
                        <th class="px-8 py-5">Dernière Collecte</th>
                    </tr>
                </thead>
                <tbody class="divide-y divide-gray-50">
                    <?php foreach ($boxes as $box): ?>
                    <tr class="hover:bg-ucu-grey/40 transition-all group">
                        <td class="px-8 py-6 font-mono font-bold text-ucu-foret"><?= $box['matricule'] ?></td>
                        <td class="px-8 py-6">
                            <p class="font-bold"><?= $box['localisation'] ?></p>
                            <p class="text-[9px] text-gray-400 italic uppercase"><?= $box['ville'] ?></p>
                        </td>
                        <td class="px-8 py-6">
                            <?php $color = ($box['occupation'] > 80) ? 'red-500' : 'ucu-lagon'; ?>
                            <div class="w-24 bg-gray-100 h-1.5 rounded-full overflow-hidden">
                                <div class="bg-<?= $color ?> h-full" style="width: <?= $box['occupation'] ?>%"></div>
                            </div>
                            <p class="text-[8px] mt-1 font-bold text-<?= $color ?> uppercase"><?= $box['occupation'] ?>% Rempli</p>
                        </td>
                        <td class="px-8 py-6 text-gray-400"><?= $box['derniere_collecte'] ?></td>
                    </tr>
                    <?php endforeach; ?>
                </tbody>
            </table>
        </section>
    </main>
</body>
</html>