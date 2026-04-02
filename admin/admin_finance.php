<?php
// On récupère les données financières depuis l'API Go
$data = file_get_contents("http://localhost:8080/api/admin/finance");
$finance = json_decode($data, true);

// Fallback (sécurité si l'API ne répond pas)
if (!$finance) {
    $finance = [
        "ca_total" => 0,
        "charges_totales" => 0,
        "resultat_net" => 0,
        "taux_benefice" => 0,
        "transactions" => []
    ];
}
?>
<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>UCU Admin | Gestion Financière</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600&family=Montserrat:wght@700&display=swap" rel="stylesheet">

    <script>
        tailwind.config = {
            theme: {
                extend: {
                    colors: {
                        'ucu-foret': '#1E4E4E',
                        'ucu-lagon': '#6A9A9A',
                        'ucu-grey': '#F8FAFC',
                    },
                    fontFamily: {
                        'sans': ['Inter', 'sans-serif'],
                        'montserrat': ['Montserrat', 'sans-serif'],
                    },
                    borderRadius: { 'ucu': '8px' }
                }
            }
        }
    </script>
</head>
<body class="bg-ucu-grey font-sans text-gray-800">

    <?php include '../includes/sidebar.php'; ?>

    <main class="lg:ml-72 p-12">
        <header class="mb-12">
            <h2 class="font-montserrat font-bold text-3xl text-ucu-foret uppercase tracking-tighter">Suivi <span class="text-ucu-lagon italic">Financier</span></h2>
            <p class="text-sm text-gray-400 mt-2">Analyse des flux Stripe, abonnements professionnels et charges d'infrastructure.</p>
        </header>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-8 mb-12">
            <div class="bg-white p-8 rounded-ucu border border-gray-100 shadow-sm">
                <p class="text-[9px] font-bold text-gray-400 uppercase mb-2">Chiffre d'Affaires (CA)</p>
                <p class="text-3xl font-montserrat font-bold text-ucu-foret"><?= number_format($finance['ca_total'], 2, ',', ' ') ?> €</p>
                <div class="mt-4 h-1 bg-gray-100 rounded-full overflow-hidden">
                    <div class="bg-green-500 h-full w-[75%]"></div>
                </div>
            </div>

            <div class="bg-white p-8 rounded-ucu border border-gray-100 shadow-sm">
                <p class="text-[9px] font-bold text-gray-400 uppercase mb-2">Total des Charges</p>
                <p class="text-3xl font-montserrat font-bold text-red-500"><?= number_format($finance['charges_totales'], 2, ',', ' ') ?> €</p>
                <p class="text-[10px] text-gray-400 mt-2 italic">Salaires, Cloud, Maintenance</p>
            </div>

            <div class="bg-white p-8 rounded-ucu border border-gray-100 shadow-sm">
                <p class="text-[9px] font-bold text-gray-400 uppercase mb-2">Résultat Net</p>
                <p class="text-3xl font-montserrat font-bold text-ucu-lagon"><?= number_format($finance['resultat_net'], 2, ',', ' ') ?> €</p>
                <p class="text-[10px] <?= $finance['taux_benefice'] >= 0 ? 'text-green-500' : 'text-red-500' ?> font-bold mt-2">
                    Bénéfice : <?= $finance['taux_benefice'] > 0 ? '+' : '' ?><?= $finance['taux_benefice'] ?>%
                </p>
            </div>
        </div>

        <section class="bg-white rounded-ucu shadow-sm border border-gray-100 overflow-hidden">
            <div class="p-6 border-b border-gray-100 flex justify-between items-center bg-ucu-grey/20">
                <h3 class="font-montserrat font-bold text-[10px] uppercase text-ucu-foret tracking-widest">Derniers flux de trésorerie</h3>
                <button class="text-[10px] font-bold text-ucu-lagon uppercase hover:underline">Exporter PDF (Stripe)</button>
            </div>
            <table class="w-full text-left text-[11px]">
                <thead class="text-gray-400 font-montserrat uppercase text-[8px] tracking-[0.2em] border-b border-gray-50">
                    <tr>
                        <th class="px-8 py-5">Date</th>
                        <th class="px-8 py-5">Désignation</th>
                        <th class="px-8 py-5">Type</th>
                        <th class="px-8 py-5">Montant</th>
                    </tr>
                </thead>
                <tbody class="divide-y divide-gray-50">
                    <?php if (!empty($finance['transactions'])): ?>
                        <?php foreach ($finance['transactions'] as $trans): ?>
                        <tr class="hover:bg-ucu-grey/40 transition-all">
                            <td class="px-8 py-4"><?= $trans['date'] ?></td>
                            <td class="px-8 py-4 font-bold text-ucu-foret"><?= $trans['designation'] ?></td>
                            <td class="px-8 py-4">
                                <span class="<?= $trans['type'] === 'Revenu' ? 'text-green-500' : 'text-red-500' ?> font-bold uppercase text-[9px]">
                                    <?= $trans['type'] ?>
                                </span>
                            </td>
                            <td class="px-8 py-4 font-bold <?= $trans['type'] === 'Revenu' ? '' : 'text-red-500' ?>">
                                <?= $trans['type'] === 'Revenu' ? '+ ' : '- ' ?><?= number_format($trans['montant'], 2, ',', ' ') ?> €
                            </td>
                        </tr>
                        <?php endforeach; ?>
                    <?php else: ?>
                        <tr>
                            <td colspan="4" class="px-8 py-10 text-center text-gray-400 italic text-[10px]">Aucune transaction récente récupérée via l'API.</td>
                        </tr>
                    <?php endif; ?>
                </tbody>
            </table>
        </section>
    </main>
</body>
</html>