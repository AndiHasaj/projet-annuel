<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>UCU | Mes Annonces</title>
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

    <nav class="bg-white border-b border-gray-100 sticky top-0 z-50">
        <div class="max-w-7xl mx-auto px-6 h-20 flex items-center justify-between">
            <div class="flex items-center gap-3">
                <a href="particulier_dashboard.php" class="w-10 h-10 bg-ucu-foret rounded-ucu flex items-center justify-center text-white font-montserrat font-bold italic hover:bg-ucu-lagon transition-all">UCU</a>
                <h1 class="font-montserrat font-bold text-ucu-foret text-sm uppercase tracking-widest">Mes Annonces</h1>
            </div>
            <div class="flex gap-6 text-[10px] font-bold uppercase tracking-widest">
                <a href="particulier_dashboard.php" class="text-gray-400 hover:text-ucu-foret">Dashboard</a>
                <a href="particulier_catalogue.php" class="text-gray-400 hover:text-ucu-foret">Catalogue</a>
                <a href="particulier_conseils.php" class="text-gray-400 hover:text-ucu-foret">Conseils</a>
                <a href="particulier_annonces.php" class="text-ucu-foret border-b-2 border-ucu-lagon pb-1">Annonces</a>
                <a href="particulier_conteneurs.php" class="text-gray-400 hover:text-ucu-foret">Conteneurs</a>
                <a href="particulier_planning.php" class="text-gray-400 hover:text-ucu-foret">Planning</a>
                <a href="particulier_profil.php" class="text-gray-400 hover:text-ucu-foret">Profil</a>
            </div>
        </div>
    </nav>

    <main class="max-w-6xl mx-auto px-6 py-12">
        
        <div class="grid grid-cols-1 lg:grid-cols-12 gap-12">
            
            <section class="lg:col-span-5 bg-white p-8 rounded-ucu shadow-sm border border-gray-100">
                <h2 class="font-montserrat font-bold text-lg text-ucu-foret uppercase mb-6 tracking-tight italic">Déposer un objet</h2>
                
                <form class="space-y-6">
                    <div class="grid grid-cols-2 gap-3">
                        <label class="cursor-pointer group">
                            <input type="radio" name="type" value="don" class="hidden peer" checked>
                            <div class="p-4 border-2 border-ucu-grey rounded-ucu text-center peer-checked:border-ucu-lagon peer-checked:bg-ucu-lagon/5 transition-all">
                                <p class="text-[10px] font-bold uppercase text-gray-400 peer-checked:text-ucu-lagon">🎁 Don</p>
                            </div>
                        </label>
                        <label class="cursor-pointer group">
                            <input type="radio" name="type" value="vente" class="hidden peer">
                            <div class="p-4 border-2 border-ucu-grey rounded-ucu text-center peer-checked:border-ucu-lagon peer-checked:bg-ucu-lagon/5 transition-all">
                                <p class="text-[10px] font-bold uppercase text-gray-400 peer-checked:text-ucu-lagon">💰 Vente</p>
                            </div>
                        </label>
                    </div>

                    <div>
                        <label class="text-[9px] font-bold uppercase text-ucu-foret mb-2 block tracking-widest">Nom de l'objet</label>
                        <input type="text" placeholder="Ex: Table basse en chêne" class="w-full bg-ucu-grey p-4 rounded-ucu text-xs outline-none focus:ring-2 focus:ring-ucu-lagon/30 border border-transparent focus:border-ucu-lagon">
                    </div>

                    <div class="grid grid-cols-2 gap-4">
                        <div>
                            <label class="text-[9px] font-bold uppercase text-ucu-foret mb-2 block tracking-widest">Catégorie</label>
                            <select class="w-full bg-ucu-grey p-4 rounded-ucu text-xs outline-none border border-transparent">
                                <option>Ameublement</option>
                                <option>Textile</option>
                                <option>Matériaux bruts</option>
                            </select>
                        </div>
                        <div>
                            <label class="text-[9px] font-bold uppercase text-ucu-foret mb-2 block tracking-widest">Prix (si vente)</label>
                            <input type="number" placeholder="0.00 €" class="w-full bg-ucu-grey p-4 rounded-ucu text-xs outline-none">
                        </div>
                    </div>

                    <div class="border-2 border-dashed border-ucu-grey p-8 rounded-ucu text-center hover:border-ucu-lagon transition-all cursor-pointer">
                        <span class="text-2xl block mb-2">📸</span>
                        <p class="text-[9px] font-bold text-gray-400 uppercase">Ajouter une photo</p>
                        <p class="text-[8px] text-gray-300 italic mt-1">(Requis pour le service Check)</p>
                    </div>

                    <button type="submit" class="w-full bg-ucu-foret text-white py-4 rounded-ucu font-montserrat font-bold text-[10px] uppercase tracking-[0.2em] shadow-lg hover:bg-ucu-lagon transition-all">
                        Envoyer pour validation
                    </button>
                </form>
            </section>

            <section class="lg:col-span-7 space-y-6">
                <h2 class="font-montserrat font-bold text-lg text-ucu-foret uppercase tracking-tight italic">Suivi de mes objets</h2>
                
                <div class="bg-white rounded-ucu shadow-sm border border-gray-100 overflow-hidden">
                    <table class="w-full text-left text-[11px]">
                        <thead>
                            <tr class="bg-ucu-grey text-gray-400 font-montserrat uppercase text-[9px] tracking-widest border-b">
                                <th class="px-6 py-5">Objet</th>
                                <th class="px-6 py-5">Mode</th>
                                <th class="px-6 py-5">Statut</th>
                                <th class="px-6 py-5 text-right">Action</th>
                            </tr>
                        </thead>
                        <tbody class="divide-y divide-gray-50">
                            <tr class="hover:bg-ucu-grey transition-all group">
                                <td class="px-6 py-5">
                                    <p class="font-bold text-ucu-foret uppercase">Chaises Vintage (x4)</p>
                                    <p class="text-[9px] text-gray-400 italic">Posté le 14/03/2026</p>
                                </td>
                                <td class="px-6 py-5 font-bold text-ucu-lagon uppercase">Vente (60€)</td>
                                <td class="px-6 py-5">
                                    <span class="bg-orange-50 text-orange-400 px-3 py-1 rounded-full text-[9px] font-bold uppercase">En cours de validation</span>
                                </td>
                                <td class="px-6 py-5 text-right">
                                    <button class="text-gray-300 hover:text-red-500 font-bold uppercase text-[9px]">Supprimer</button>
                                </td>
                            </tr>
                            <tr class="hover:bg-ucu-grey transition-all group">
                                <td class="px-6 py-5">
                                    <p class="font-bold text-ucu-foret uppercase">Étagère Industrielle</p>
                                    <p class="text-[9px] text-gray-400 italic">Posté le 10/03/2026</p>
                                </td>
                                <td class="px-6 py-5 font-bold text-ucu-lagon uppercase">Don</td>
                                <td class="px-6 py-5">
                                    <span class="bg-green-50 text-green-600 px-3 py-1 rounded-full text-[9px] font-bold uppercase">Validée</span>
                                </td>
                                <td class="px-6 py-5 text-right">
                                    <button class="text-ucu-foret hover:underline font-bold uppercase text-[9px]">Voir l'annonce</button>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>

                <div class="bg-ucu-foret p-6 rounded-ucu text-white flex gap-4 items-center shadow-lg">
                    <div class="text-2xl">💡</div>
                    <p class="text-[10px] leading-relaxed italic">
                        <span class="text-ucu-lagon font-bold uppercase">Important :</span> Une fois votre annonce validée, vous recevrez une notification OneSignal. Si c'est un don, l'artisan pourra scanner votre code barre pour la récupération.
                    </p>
                </div>
            </section>

        </div>
    </main>

</body>
</html>