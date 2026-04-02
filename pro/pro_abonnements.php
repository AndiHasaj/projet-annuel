<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>UCU Pro | Abonnements & Factures</title>
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

    <nav class="bg-ucu-foret border-b border-white/10 sticky top-0 z-50 text-white">
        <div class="max-w-7xl mx-auto px-6 h-20 flex items-center justify-between">
            <div class="flex items-center gap-3">
                <a href="pro_dashboard.php" class="w-10 h-10 bg-white rounded-ucu flex items-center justify-center text-ucu-foret font-montserrat font-bold italic">PRO</a>
                <h1 class="font-montserrat font-bold text-sm uppercase tracking-widest">Contrats & Finances</h1>
            </div>
            <div class="flex gap-6 text-[10px] font-bold uppercase tracking-widest text-white/60">
                <a href="pro_dashboard.php" class="hover:text-white">Dashboard</a>
                <a href="pro_logistique.php" class="hover:text-white">Logistique</a>
                <a href="pro_abonnements.php" class="text-ucu-lagon">Abonnement</a>
            </div>
        </div>
    </nav>

    <main class="max-w-6xl mx-auto px-6 py-12">
        
        <div class="grid grid-cols-1 lg:grid-cols-12 gap-12">
            
            <section class="lg:col-span-7 space-y-8">
                <div class="bg-white p-8 rounded-ucu shadow-sm border border-gray-100">
                    <div class="flex justify-between items-start mb-8">
                        <div>
                            <h2 class="font-montserrat font-bold text-xl text-ucu-foret uppercase italic">Contrat Pro Premium</h2>
                            <p class="text-[10px] text-gray-400 font-bold uppercase tracking-widest mt-1">ID : UCU-SUBS-2026-AF89</p>
                        </div>
                        <span class="bg-green-50 text-green-600 px-4 py-1 rounded-full text-[9px] font-bold uppercase">Actif</span>
                    </div>

                    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
                        <div class="p-6 bg-ucu-grey rounded-ucu border border-gray-100">
                            <p class="text-[9px] font-bold text-gray-400 uppercase mb-2">Mensualité</p>
                            <p class="text-2xl font-montserrat font-bold text-ucu-foret">49.90 €<span class="text-xs opacity-40">/mois</span></p>
                        </div>
                        <div class="p-6 bg-ucu-grey rounded-ucu border border-gray-100">
                            <p class="text-[9px] font-bold text-gray-400 uppercase mb-2">Prochain prélèvement</p>
                            <p class="text-2xl font-montserrat font-bold text-ucu-lagon">01 Avril</p>
                        </div>
                    </div>

                    <div class="space-y-4">
                        <h3 class="font-montserrat font-bold text-[10px] uppercase text-ucu-foret tracking-widest border-b pb-2">Options incluses</h3>
                        <ul class="text-[11px] space-y-3 italic text-gray-500">
                            <li class="flex items-center gap-2">✅ Accès illimité aux annonces de dons</li>
                            <li class="flex items-center gap-2">✅ Mise en avant de 3 projets / mois</li>
                            <li class="flex items-center gap-2">✅ Badge "Artisan Vérifié" sur le catalogue</li>
                        </ul>
                    </div>
                </div>

                <div class="bg-white rounded-ucu shadow-sm border border-gray-100 overflow-hidden">
                    <div class="p-6 border-b border-gray-100 bg-ucu-grey/20">
                        <h3 class="font-montserrat font-bold text-[10px] uppercase text-ucu-foret tracking-widest">Historique des factures</h3>
                    </div>
                    <table class="w-full text-left text-[11px]">
                        <thead>
                            <tr class="text-gray-400 font-montserrat uppercase text-[9px] tracking-widest border-b">
                                <th class="px-6 py-4">Période</th>
                                <th class="px-6 py-4">Montant</th>
                                <th class="px-6 py-4">Status</th>
                                <th class="px-6 py-4 text-right">PDF</th>
                            </tr>
                        </thead>
                        <tbody class="divide-y divide-gray-50">
                            <tr class="hover:bg-ucu-grey/30 transition-all">
                                <td class="px-6 py-5 font-bold uppercase">Mars 2026</td>
                                <td class="px-6 py-5 font-mono">49.90 €</td>
                                <td class="px-6 py-5"><span class="text-green-600 font-bold uppercase text-[9px]">Payé</span></td>
                                <td class="px-6 py-5 text-right">
                                    <button class="text-ucu-foret hover:text-ucu-lagon transition-colors">📥 Télécharger</button>
                                </td>
                            </tr>
                            <tr class="hover:bg-ucu-grey/30 transition-all">
                                <td class="px-6 py-5 font-bold uppercase">Février 2026</td>
                                <td class="px-6 py-5 font-mono">49.90 €</td>
                                <td class="px-6 py-5"><span class="text-green-600 font-bold uppercase text-[9px]">Payé</span></td>
                                <td class="px-6 py-5 text-right">
                                    <button class="text-ucu-foret hover:text-ucu-lagon transition-colors">📥 Télécharger</button>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </section>

            <aside class="lg:col-span-5 space-y-8">
                <div class="bg-white p-8 rounded-ucu shadow-sm border border-gray-100">
                    <h3 class="font-montserrat font-bold text-[10px] uppercase text-ucu-foret mb-6 tracking-widest">Moyen de paiement</h3>
                    <div class="p-5 border-2 border-ucu-lagon/30 rounded-ucu flex items-center justify-between mb-6">
                        <div class="flex items-center gap-4">
                            <div class="text-2xl">💳</div>
                            <div>
                                <p class="text-xs font-bold text-ucu-foret">•••• •••• •••• 4242</p>
                                <p class="text-[9px] text-gray-400 uppercase">Expire : 12/28</p>
                            </div>
                        </div>
                        <button class="text-[9px] font-bold text-ucu-lagon uppercase hover:underline">Editer</button>
                    </div>
                    <p class="text-[9px] text-gray-400 italic text-center">Sécurisé par l'infrastructure **Stripe Connect**</p>
                </div>

                <div class="bg-ucu-foret p-8 rounded-ucu shadow-xl text-white relative overflow-hidden">
                    <h3 class="font-montserrat font-bold text-xs uppercase mb-4 tracking-tighter">Boostez votre visibilité</h3>
                    <p class="text-[10px] opacity-70 mb-6 leading-relaxed italic">
                        Apparaissez en tête des résultats du catalogue particulier pendant 7 jours.
                    </p>
                    <div class="bg-white/10 p-4 rounded-ucu mb-6 flex justify-between items-center">
                        <span class="text-[10px] font-bold uppercase tracking-widest">Option Pub</span>
                        <span class="text-sm font-montserrat font-bold text-ucu-lagon">15.00 €</span>
                    </div>
                    <button class="w-full bg-ucu-lagon text-white py-3 rounded-ucu font-bold text-[9px] uppercase tracking-widest hover:bg-white hover:text-ucu-foret transition-all">
                        Activer maintenant
                    </button>
                    <div class="absolute -right-4 -bottom-4 text-6xl opacity-10 rotate-12">🚀</div>
                </div>

                <div class="text-center">
                    <button class="text-[9px] font-bold text-gray-300 uppercase hover:text-red-500 transition-colors">Suspendre mon contrat</button>
                </div>
            </aside>

        </div>
    </main>

    <footer class="mt-20 py-10 border-t border-gray-100 text-center">
        <p class="text-[9px] font-bold text-gray-400 uppercase tracking-[0.3em]">UpcycleConnect Pro • Facturation & Fiscalité</p>
    </footer>

</body>
</html>