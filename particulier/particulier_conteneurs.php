<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>UCU | Dépôts en Conteneur</title>
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
                <a href="particulier_dashboard.php" class="w-10 h-10 bg-ucu-foret rounded-ucu flex items-center justify-center text-white font-montserrat font-bold italic">UCU</a>
                <h1 class="font-montserrat font-bold text-ucu-foret text-sm uppercase tracking-widest">Logistique Dépôts</h1>
            </div>
            <div class="flex gap-6 text-[10px] font-bold uppercase tracking-widest">
                <a href="particulier_dashboard.php" class="text-gray-400 hover:text-ucu-foret">Dashboard</a>
                <a href="particulier_catalogue.php" class="text-gray-400 hover:text-ucu-foret">Catalogue</a>
                <a href="particulier_conseils.php" class="text-gray-400 hover:text-ucu-foret">Conseils</a>
                <a href="particulier_annonces.php" class="text-gray-400 hover:text-ucu-foret">Annonces</a>
                <a href="particulier_conteneurs.php" class="text-ucu-foret border-b-2 border-ucu-lagon pb-1">Conteneurs</a>
                <a href="particulier_planning.php" class="text-gray-400 hover:text-ucu-foret">Planning</a>
                <a href="particulier_profil.php" class="text-gray-400 hover:text-ucu-foret">Profil</a>
            </div>
        </div>
    </nav>

    <main class="max-w-6xl mx-auto px-6 py-12">
        
        <header class="mb-12">
            <h2 class="font-montserrat font-bold text-3xl text-ucu-foret uppercase tracking-tighter">Accès aux <span class="text-ucu-lagon italic">Conteneurs</span></h2>
            <p class="text-sm text-gray-400 mt-2">Demandez un code pour déposer vos objets encombrants en toute sécurité.</p>
        </header>

        <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
            
            <div class="lg:col-span-1">
                <section class="bg-white p-8 rounded-ucu shadow-sm border border-gray-100">
                    <h3 class="font-montserrat font-bold text-xs uppercase text-ucu-foret mb-6 tracking-widest border-b border-ucu-grey pb-4">Nouvelle Demande</h3>
                    
                    <form class="space-y-6">
                        <div>
                            <label class="text-[9px] font-bold uppercase text-gray-400 mb-2 block">Choisir la Box</label>
                            <select class="w-full bg-ucu-grey p-4 rounded-ucu text-xs outline-none border border-transparent focus:border-ucu-lagon">
                                <option>Box Paris 10e - La Fayette</option>
                                <option>Box Ivry-sur-Seine - Centre</option>
                                <option>Box Montreuil - Mairie</option>
                            </select>
                        </div>

                        <div>
                            <label class="text-[9px] font-bold uppercase text-gray-400 mb-2 block">Description de l'objet</label>
                            <textarea placeholder="Dimensions, poids approximatif..." class="w-full bg-ucu-grey p-4 rounded-ucu text-xs outline-none h-24 resize-none"></textarea>
                        </div>

                        <div class="border-2 border-dashed border-ucu-grey p-6 rounded-ucu text-center">
                            <p class="text-[9px] font-bold text-gray-400 uppercase">Photo de l'objet</p>
                            <p class="text-[8px] text-gray-300 italic mt-1">Nécessaire pour le contrôle Check</p>
                        </div>

                        <button type="submit" class="w-full bg-ucu-foret text-white py-4 rounded-ucu font-montserrat font-bold text-[10px] uppercase tracking-widest hover:bg-ucu-lagon transition-all shadow-md">
                            Solliciter un code
                        </button>
                    </form>
                </section>
            </div>

            <div class="lg:col-span-2 space-y-8">
                
                <div class="bg-ucu-foret p-10 rounded-ucu shadow-xl text-white relative overflow-hidden">
                    <div class="relative z-10 flex flex-col md:flex-row justify-between items-center gap-8">
                        <div>
                            <span class="bg-ucu-lagon text-ucu-foret px-3 py-1 rounded-full text-[9px] font-bold uppercase mb-4 inline-block">Demande Approuvée</span>
                            <h3 class="font-montserrat font-bold text-2xl uppercase tracking-tighter">Votre Code d'accès</h3>
                            <p class="text-xs opacity-60 mt-1 italic">Valable 48h à la Box Ivry-Centre</p>
                        </div>
                        <div class="bg-white text-ucu-foret p-6 rounded-ucu text-center min-w-[180px]">
                            <p class="text-[9px] font-bold uppercase text-gray-400 mb-2">Code Unique</p>
                            <p class="text-4xl font-montserrat font-bold tracking-[0.2em]">8492</p>
                        </div>
                    </div>
                    <div class="absolute -right-12 -bottom-12 w-48 h-48 bg-white/5 rounded-full"></div>
                </div>

                <section class="bg-white rounded-ucu shadow-sm border border-gray-100 overflow-hidden">
                    <div class="p-6 border-b border-gray-100 flex justify-between items-center">
                        <h4 class="font-montserrat font-bold text-[10px] uppercase text-ucu-foret tracking-widest italic">Mes dépôts en cours</h4>
                    </div>
                    <table class="w-full text-left text-[11px]">
                        <thead class="bg-ucu-grey text-gray-400 font-montserrat uppercase text-[9px] tracking-widest">
                            <tr>
                                <th class="px-6 py-4 italic">Objet</th>
                                <th class="px-6 py-4">Box</th>
                                <th class="px-6 py-4">Status</th>
                                <th class="px-6 py-4 text-right">Récupération</th>
                            </tr>
                        </thead>
                        <tbody class="divide-y divide-gray-50">
                            <tr class="hover:bg-ucu-grey transition-all group">
                                <td class="px-6 py-5 font-bold uppercase text-ucu-foret">Lot de palettes bois</td>
                                <td class="px-6 py-5 italic text-gray-500">Paris 10e</td>
                                <td class="px-6 py-5">
                                    <span class="text-orange-400 font-bold uppercase">Vérification Photo...</span>
                                </td>
                                <td class="px-6 py-5 text-right">
                                    <span class="text-gray-300 italic">En attente</span>
                                </td>
                            </tr>
                            <tr class="hover:bg-ucu-grey transition-all group">
                                <td class="px-6 py-5 font-bold uppercase text-ucu-foret">Chute de métal (Cuivre)</td>
                                <td class="px-6 py-5 italic text-gray-500">Ivry-Centre</td>
                                <td class="px-6 py-5 text-green-600 font-bold uppercase">Déposé</td>
                                <td class="px-6 py-5 text-right">
                                    <button class="bg-ucu-lagon/10 text-ucu-lagon px-3 py-1.5 rounded-ucu font-bold uppercase text-[9px] hover:bg-ucu-lagon hover:text-white transition-all">
                                        Voir Code Barre Pro
                                    </button>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </section>

                <div class="bg-white p-6 rounded-ucu border-l-4 border-ucu-foret shadow-sm">
                    <div class="flex gap-4">
                        <span class="text-xl">ℹ️</span>
                        <div>
                            <p class="font-bold text-xs text-ucu-foret uppercase mb-1">Comment ça marche ?</p>
                            <p class="text-[10px] text-gray-500 leading-relaxed italic">
                                Une fois votre objet déposé, nous générons un **Code Barre** unique. Lorsqu'un artisan récupère votre objet, votre **Upcycling Score** est automatiquement crédité.
                            </p>
                        </div>
                    </div>
                </div>
            </div>

        </div>
    </main>

</body>
</html>