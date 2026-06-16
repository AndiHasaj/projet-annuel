<?php
$response_particuliers = file_get_contents("http://localhost:8080/particuliers/");
$response_salaries = file_get_contents("http://localhost:8080/salaries/");
$response_professionnels = file_get_contents("http://localhost:8080/professionnels/");
$particuliers = json_decode($response_particuliers, true) ?? [];
$salaries = json_decode($response_salaries, true) ?? [];
$professionnels = json_decode($response_professionnels, true) ?? [];

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
                        <th class="px-8 py-5">Particuliers</th>
                        <th class="px-8 py-5">Score</th>
                        <th class="px-8 py-5">Dernière Connexion</th>
                    </tr>
                </thead>
                <tbody class="divide-y divide-gray-50">
                    <?php foreach ($particuliers as $particulier): ?>
                    <tr class="hover:bg-ucu-grey/40 transition-all group">
                        <td class="px-8 py-6">
                            <p class="font-bold text-ucu-foret"><?= $particulier['nom'] . ' ' . $particulier['prenom'] ?></p>
                            <p class="text-[9px] text-gray-400 uppercase"><?= $particulier['email'] ?></p>
                        </td>
                        <td class="px-8 py-6">
                            <span class="text-[9px] font-bold text-ucu-lagon uppercase"><?= $particulier['score'] ?></span>
                        </td>
                        <td class="px-8 py-6">
                            <span class="text-[9px] font-bold text-ucu-lagon uppercase"><?= $particulier['derniere_connexion'] ?></span>
                        </td>
                    </tr>
                    <?php endforeach; ?>
                </tbody>
            </table>
        </section>
        <section class="bg-white rounded-ucu shadow-sm border border-gray-100 overflow-hidden">
            <table class="w-full text-left text-[11px]">
                <thead class="bg-ucu-grey border-b border-gray-100 text-gray-400 font-montserrat uppercase text-[9px] tracking-widest">
                    <tr>
                        <th class="px-8 py-5">Professionnels</th>
                        <th class="px-8 py-5">Site Web</th>
                        <th class="px-8 py-5">Numéro de Siret</th>
                        <th class="px-8 py-5">Derniére connexion</th>
                    </tr>
                </thead>
                <tbody class="divide-y divide-gray-50">
                    <?php foreach ($professionnels as $professionnel): ?>
                    <tr class="hover:bg-ucu-grey/40 transition-all group">
                        <td class="px-8 py-6">
                            <p class="font-bold text-ucu-foret"><?= $professionnel['nom_entreprise'] ?></p>
                            <p class="text-[9px] text-gray-400 uppercase"><?= $professionnel['email'] ?></p>
                        </td>
                        <td class="px-8 py-6">
                            <span class="text-[9px] font-bold text-ucu-lagon uppercase"><?= $professionnel['site_web'] ?></span>
                        </td>
                        <td class="px-8 py-6">
                            <span class="text-[9px] font-bold text-ucu-lagon uppercase"><?= $professionnel['numero_siret'] ?></span>
                        </td>
                        <td class="px-8 py-6 text-gray-400 italic">
                            <?= $professionnel['derniere_connexion'] ?>
                        </td>
                    </tr>
                    <?php endforeach; ?>
                </tbody>
            </table>
        </section>
        <section class="bg-white rounded-ucu shadow-sm border border-gray-100 overflow-hidden">
            <table class="w-full text-left text-[11px]">
                <thead class="bg-ucu-grey border-b border-gray-100 text-gray-400 font-montserrat uppercase text-[9px] tracking-widest">
                    <tr>
                        <th class="px-8 py-5">Salariés</th>
                        <th class="px-8 py-5">Poste</th>
                        <th class="px-8 py-5">Contrat</th>
                        <th class="px-8 py-5">Salaire</th>
                        <th class="px-8 py-5">Dernière Connexion</th>
                    </tr>
                </thead>
                <tbody class="divide-y divide-gray-50">
                    <?php foreach ($salaries as $salarie): ?>
                    <tr class="hover:bg-ucu-grey/40 transition-all group">
                        <td class="px-8 py-6">
                            <p class="font-bold text-ucu-foret"><?= $salarie['nom'] . ' ' . $salarie['prenom'] ?></p>
                            <p class="text-[9px] text-gray-400 uppercase"><?= $salarie['email'] ?></p>
                        </td>
                        <td class="px-8 py-6">
                            <span class="text-[9px] font-bold text-ucu-lagon uppercase"><?= $salarie['poste'] ?></span>
                        </td>
                        <td class="px-8 py-6 text-gray-400 italic">
                            <span class="text-[9px] font-bold text-ucu-lagon uppercase"><?= $salarie['contrat'] ?></span>
                        </td>
                        <td class="px-8 py-6">
                            <span class="text-[9px] font-bold text-ucu-lagon uppercase"><?= $salarie['salaire'] ?></span>
                        </td>
                        <td class="px-8 py-6 text-gray-400 italic">
                            <?= $salarie['derniere_connexion'] ?>
                        </td>
                    </tr>
                    <?php endforeach; ?>
                </tbody>
            </table>
        </section>
    </main>
</body>
</html>