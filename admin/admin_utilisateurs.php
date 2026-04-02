<?php
$response = file_get_contents("http://localhost:8080/api/admin/users");
$users = json_decode($response, true) ?? [];
?>
<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <title>UCU Admin | Utilisateurs</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600&family=Montserrat:wght@700&display=swap" rel="stylesheet">
    <script>
        tailwind.config = { theme: { extend: { colors: { 'ucu-foret': '#1E4E4E', 'ucu-lagon': '#6A9A9A', 'ucu-grey': '#F8FAFC' }, fontFamily: { 'sans': ['Inter', 'sans-serif'], 'montserrat': ['Montserrat', 'sans-serif'] }, borderRadius: { 'ucu': '8px' } } } }
    </script>
</head>
<body class="bg-ucu-grey font-sans text-gray-800">
    <?php include '../includes/sidebar.php'; ?>
    <main class="lg:ml-72 p-12">
        <header class="mb-12">
            <h2 class="font-montserrat font-bold text-3xl text-ucu-foret uppercase tracking-tighter">Annuaire <span class="text-ucu-lagon italic">Acteurs</span></h2>
        </header>

        <section class="bg-white rounded-ucu shadow-sm border border-gray-100 overflow-hidden">
            <table class="w-full text-left text-[11px]">
                <thead class="bg-ucu-grey border-b border-gray-100 text-gray-400 font-montserrat uppercase text-[9px] tracking-widest">
                    <tr>
                        <th class="px-8 py-5">Utilisateur</th>
                        <th class="px-8 py-5">Rôle</th>
                        <th class="px-8 py-5">Détail</th>
                        <th class="px-8 py-5">Status</th>
                    </tr>
                </thead>
                <tbody class="divide-y divide-gray-50">
                    <?php foreach ($users as $user): ?>
                    <tr class="hover:bg-ucu-grey/40 transition-all group">
                        <td class="px-8 py-6">
                            <p class="font-bold text-ucu-foret"><?= $user['nom'] ?></p>
                            <p class="text-[9px] text-gray-400 uppercase"><?= $user['email'] ?></p>
                        </td>
                        <td class="px-8 py-6">
                            <span class="text-[9px] font-bold text-ucu-lagon uppercase"><?= $user['role'] ?></span>
                        </td>
                        <td class="px-8 py-6 text-gray-400 italic">
                            <?= $user['detail'] ?>
                        </td>
                        <td class="px-8 py-6">
                            <?php $statusColor = ($user['status'] == 'Actif') ? 'green-100 text-green-600' : 'orange-100 text-orange-500'; ?>
                            <span class="bg-<?= $statusColor ?> px-2 py-0.5 rounded text-[8px] font-bold uppercase"><?= $user['status'] ?></span>
                        </td>
                    </tr>
                    <?php endforeach; ?>
                </tbody>
            </table>
        </section>
    </main>
</body>
</html>