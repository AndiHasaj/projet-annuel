<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>UCU Pro | Espace Artisan</title>
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
                <div class="w-10 h-10 bg-white rounded-ucu flex items-center justify-center text-ucu-foret font-montserrat font-bold italic">PRO</div>
                <h1 class="font-montserrat font-bold text-sm uppercase tracking-widest">Atelier Artisan</h1>
            </div>
            <div class="flex gap-8 text-[10px] font-bold uppercase tracking-widest">
                <a href="pro_dashboard.php" class="text-ucu-lagon">Dashboard</a>
                <a href="pro_marche.php" class="hover:text-ucu-lagon">Marché Matières</a>
                <a href="pro_logistique.php" class="hover:text-ucu-lagon">Récupération</a>
            </div>
        </div>
    </nav>

    <main class="max-w-7xl mx-auto px-6 py-12">
        
        <div class="grid grid-cols-1 lg:grid-cols-12 gap-8">
            
            <div class="lg:col-span-4 space-y-6">
                <section class="bg-white p-8 rounded-ucu shadow-sm border border-gray-100">
                    <h3 class="font-montserrat font-bold text-[10px] uppercase text-ucu-foret mb-6 tracking-widest">Mon Contrat</h3>
                    <div class="bg-ucu-grey p-5 rounded-ucu border-l-4 border-ucu-lagon mb-6">
                        <p class="text-[9px] font-bold text-gray-400 uppercase">Offre Actuelle</p>
                        <p class="text-sm font-montserrat font-bold text-ucu-foret">Pack "Artisan Premium"</p>
                        <p class="text-[10px] text-ucu-lagon mt-1">Prochaine facture : 01/04/2026</p>
                    </div>
                    <div class="space-y-3">
                        <button class="w-full bg-ucu-foret text-white py-3 rounded-ucu font-bold text-[9px] uppercase tracking-widest">Gérer mon abonnement</button>
                        <button class="w-full border border-gray-200 py-3 rounded-ucu font-bold text-[9px] uppercase tracking-widest text-gray-400 hover:bg-gray-50">Campagnes Publicitaires</button>
                    </div>
                </section>

                <section class="bg-white p-8 rounded-ucu shadow-sm border border-gray-100">
                    <h3 class="font-montserrat font-bold text-[10px] uppercase text-ucu-foret mb-6 tracking-widest">Mise en avant</h3>
                    <div class="flex items-center justify-between p-4 bg-ucu-grey rounded-ucu">
                        <p class="text-[10px] font-bold">Vues de vos projets</p>
                        <span class="text-ucu-lagon font-bold">+24%</span>
                    </div>
                </article>
            </div>

            <div class="lg:col-span-8 space-y-8">
                
                <section class="bg-white p-8 rounded-ucu shadow-sm border border-gray-100">
                    <div class="flex justify-between items-center mb-8">
                        <h3 class="font-montserrat font-bold text-sm uppercase text-ucu-foret tracking-tighter italic">Dernières opportunités de sourcing</h3>
                        <a href="pro_marche.php" class="text-[9px] font-bold text-ucu-lagon uppercase underline">Voir tout le marché</a>
                    </div>
                    
                    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                        <div class="border border-gray-100 p-4 rounded-ucu flex gap-4 hover:border-ucu-lagon transition-all cursor-pointer">
                            <div class="w-16 h-16 bg-ucu-grey rounded flex items-center justify-center text-xl">🪑</div>
                            <div>
                                <p class="text-[10px] font-bold text-ucu-foret uppercase">Lot 4 Chaises Bois</p>
                                <p class="text-[9px] text-gray-400">Ivry • Particulier (Don)</p>
                                <button class="mt-2 text-[8px] bg-ucu-lagon text-white px-2 py-1 rounded uppercase font-bold">Réserver l'objet</button>
                            </div>
                        </div>
                        <div class="border border-gray-100 p-4 rounded-ucu flex gap-4 hover:border-ucu-lagon transition-all cursor-pointer">
                            <div class="w-16 h-16 bg-ucu-grey rounded flex items-center justify-center text-xl">🪵</div>
                            <div>
                                <p class="text-[10px] font-bold text-ucu-foret uppercase">Palettes Pin Brut</p>
                                <p class="text-[9px] text-gray-400">Montreuil • 15.00€</p>
                                <button class="mt-2 text-[8px] bg-ucu-foret text-white px-2 py-1 rounded uppercase font-bold">Acheter</button>
                            </div>
                        </div>
                    </div>
                </section>

                <section class="bg-white rounded-ucu shadow-sm border border-gray-100 overflow-hidden">
                    <div class="p-6 border-b border-gray-100">
                        <h4 class="font-montserrat font-bold text-[10px] uppercase text-ucu-foret tracking-widest">Projets en cours de revalorisation</h4>
                    </div>
                    <div class="p-6">
                        <table class="w-full text-left text-[11px]">
                            <thead class="bg-ucu-grey text-gray-400 font-montserrat uppercase text-[9px]">
                                <tr>
                                    <th class="px-4 py-3">Projet</th>
                                    <th class="px-4 py-3">Matière Sourcing</th>
                                    <th class="px-4 py-3">Étape</th>
                                    <th class="px-4 py-3 text-right">Visibilité</th>
                                </tr>
                            </thead>
                            <tbody class="divide-y divide-gray-50">
                                <tr>
                                    <td class="px-4 py-4 font-bold text-ucu-foret uppercase">Buffet Scandinave Rev.</td>
                                    <td class="px-4 py-4 text-gray-400">Chêne récupéré (Box #402)</td>
                                    <td class="px-4 py-4"><span class="text-orange-400 font-bold uppercase text-[9px]">Finitions</span></td>
                                    <td class="px-4 py-4 text-right"><span class="bg-green-50 text-green-600 px-2 py-1 rounded-full font-bold">En Avant</span></td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                </section>
            </div>

        </div>
    </main>
</body>
</html>