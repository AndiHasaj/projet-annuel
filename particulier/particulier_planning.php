<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>UCU | Mon Planning</title>
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
                <h1 class="font-montserrat font-bold text-ucu-foret text-sm uppercase tracking-widest">Mon Emploi du Temps</h1>
            </div>
            <div class="flex gap-6 text-[10px] font-bold uppercase tracking-widest">
                <a href="particulier_dashboard.php" class="text-gray-400 hover:text-ucu-foret">Dashboard</a>
                <a href="particulier_catalogue.php" class="text-gray-400 hover:text-ucu-foret">Catalogue</a>
                <a href="particulier_conseils.php" class="text-gray-400 hover:text-ucu-foret">Conseils</a>
                <a href="particulier_annonces.php" class="text-gray-400 hover:text-ucu-foret">Annonces</a>
                <a href="particulier_conteneurs.php" class="text-gray-400 hover:text-ucu-foret">Conteneurs</a>
                <a href="particulier_planning.php" class="text-ucu-foret border-b-2 border-ucu-lagon pb-1">Planning</a>
                <a href="particulier_profil.php" class="text-gray-400 hover:text-ucu-foret">Profil</a>
            </div>
        </div>
    </nav>

    <main class="max-w-6xl mx-auto px-6 py-12">
        
        <header class="mb-12 flex flex-col md:flex-row justify-between items-center gap-6">
            <div>
                <h2 class="font-montserrat font-bold text-3xl text-ucu-foret uppercase tracking-tighter">Votre <span class="text-ucu-lagon italic">Agenda</span></h2>
                <p class="text-sm text-gray-400 mt-2">Retrouvez vos cours, ateliers et créneaux de dépôt.</p>
            </div>
            <div class="flex bg-white p-1 rounded-ucu shadow-sm border border-gray-100">
                <button class="px-6 py-2 bg-ucu-foret text-white rounded-ucu text-[9px] font-bold uppercase tracking-widest transition-all">Liste</button>
                <button class="px-6 py-2 text-gray-400 rounded-ucu text-[9px] font-bold uppercase tracking-widest hover:text-ucu-foret transition-all">Calendrier</button>
            </div>
        </header>

        <div class="grid grid-cols-1 lg:grid-cols-4 gap-8">
            
            <div class="lg:col-span-1 space-y-6">
                <div class="bg-ucu-foret p-8 rounded-ucu text-white shadow-lg">
                    <h4 class="font-montserrat font-bold text-[10px] uppercase text-ucu-lagon mb-6 tracking-widest">Compteur Apprentissage</h4>
                    <div class="mb-6">
                        <p class="text-4xl font-montserrat font-bold italic">12h</p>
                        <p class="text-[9px] uppercase opacity-60 font-bold tracking-widest">de formation ce mois</p>
                    </div>
                    <p class="text-[10px] opacity-80 leading-relaxed italic border-t border-white/10 pt-4">
                        "Chaque heure apprise est un pas vers une consommation plus responsable."
                    </p>
                </div>

                <div class="bg-white p-6 rounded-ucu border border-gray-100 shadow-sm text-center">
                    <p class="text-[9px] font-bold text-gray-400 uppercase mb-4 tracking-widest">Besoin d'aide ?</p>
                    <button class="text-ucu-foret font-bold text-[10px] uppercase border-b border-ucu-foret pb-1">Contacter un formateur</button>
                </div>
            </div>

            <div class="lg:col-span-3 space-y-4">
                
                <div class="relative pl-8 pb-8">
                    <div class="absolute left-0 top-0 w-px h-full bg-gray-200"></div>
                    <div class="absolute left-[-4px] top-0 w-2 h-2 rounded-full bg-ucu-lagon"></div>
                    
                    <h3 class="font-montserrat font-bold text-xs uppercase text-gray-400 tracking-[0.2em] mb-6">Mardi 17 Mars</h3>

                    <div class="bg-white p-6 rounded-ucu shadow-sm border border-gray-100 flex flex-col md:flex-row justify-between items-center gap-6 hover:border-ucu-lagon transition-all group">
                        <div class="flex gap-6 items-center">
                            <div class="text-center bg-ucu-grey p-3 rounded-ucu min-w-[80px]">
                                <p class="text-[10px] font-bold text-ucu-foret uppercase">14:00</p>
                                <p class="text-[10px] text-gray-400">16:30</p>
                            </div>
                            <div>
                                <span class="text-[8px] font-bold bg-ucu-lagon/10 text-ucu-lagon px-2 py-0.5 rounded uppercase tracking-tighter">Formation</span>
                                <h4 class="font-montserrat font-bold text-ucu-foret text-sm uppercase mt-1">Soudure sur Métaux Récupérés</h4>
                                <p class="text-[10px] text-gray-400 italic">Intervenant : Florian G. • Ivry-sur-Seine</p>
                            </div>
                        </div>
                        <div class="flex gap-3">
                            <button class="text-[9px] font-bold text-ucu-foret uppercase border border-ucu-foret px-4 py-2 rounded-ucu hover:bg-ucu-foret hover:text-white transition-all">Détails</button>
                        </div>
                    </div>
                </div>

                <div class="relative pl-8 pb-8">
                    <div class="absolute left-0 top-0 w-px h-full bg-gray-200"></div>
                    <div class="absolute left-[-4px] top-0 w-2 h-2 rounded-full bg-gray-200 group-hover:bg-ucu-lagon"></div>
                    
                    <h3 class="font-montserrat font-bold text-xs uppercase text-gray-400 tracking-[0.2em] mb-6">Jeudi 19 Mars</h3>

                    <div class="bg-white p-6 rounded-ucu shadow-sm border border-gray-100 flex flex-col md:flex-row justify-between items-center gap-6 hover:border-ucu-lagon transition-all">
                        <div class="flex gap-6 items-center">
                            <div class="text-center bg-ucu-grey p-3 rounded-ucu min-w-[80px]">
                                <p class="text-[10px] font-bold text-ucu-foret uppercase">09:30</p>
                                <p class="text-[10px] text-gray-400">10:00</p>
                            </div>
                            <div>
                                <span class="text-[8px] font-bold bg-ucu-foret/10 text-ucu-foret px-2 py-0.5 rounded uppercase tracking-tighter">Logistique</span>
                                <h4 class="font-montserrat font-bold text-ucu-foret text-sm uppercase mt-1">Dépôt Conteneur #8492</h4>
                                <p class="text-[10px] text-gray-400 italic">Lieu : Box Centre-Ville • Code : 8492</p>
                            </div>
                        </div>
                        <a href="particulier_conteneurs.php" class="text-[9px] font-bold text-ucu-lagon uppercase underline tracking-widest">Voir le code</a>
                    </div>
                </div>

                <div class="relative pl-8">
                    <div class="absolute left-0 top-0 w-px h-10 bg-gray-200"></div>
                    <div class="absolute left-[-4px] top-0 w-2 h-2 rounded-full bg-gray-200"></div>
                    
                    <h3 class="font-montserrat font-bold text-xs uppercase text-gray-400 tracking-[0.2em] mb-6">Samedi 21 Mars</h3>

                    <div class="bg-ucu-grey/50 p-6 rounded-ucu border border-dashed border-gray-200 flex flex-col md:flex-row justify-between items-center gap-6">
                        <div class="flex gap-6 items-center">
                            <div class="text-center bg-white p-3 rounded-ucu min-w-[80px]">
                                <p class="text-[10px] font-bold text-gray-400 uppercase">10:00</p>
                                <p class="text-[10px] text-gray-300 italic">Toute la journée</p>
                            </div>
                            <div>
                                <span class="text-[8px] font-bold bg-gray-100 text-gray-400 px-2 py-0.5 rounded uppercase tracking-tighter">Événement</span>
                                <h4 class="font-montserrat font-bold text-gray-400 text-sm uppercase mt-1 opacity-70">Bourse à l'Upcycling</h4>
                                <p class="text-[10px] text-gray-400 italic italic">Vous n'êtes pas encore inscrit à cet événement.</p>
                            </div>
                        </div>
                        <button class="bg-ucu-lagon text-white px-4 py-2 rounded-ucu font-bold text-[9px] uppercase hover:bg-ucu-foret transition-all">S'inscrire</button>
                    </div>
                </div>

            </div>
        </div>
    </main>

    <footer class="mt-20 py-10 border-t border-gray-100 text-center">
        <p class="text-[9px] font-bold text-gray-300 uppercase tracking-[0.3em]">UpcycleConnect • Espace Particulier • 2026</p>
    </footer>

</body>
</html>