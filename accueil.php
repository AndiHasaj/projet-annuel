<!DOCTYPE html>
<html lang="fr" class="scroll-smooth">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>UpcycleConnect | Accueil</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500&family=Montserrat:wght@700&display=swap" rel="stylesheet">
    
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
    <div id="tuto-overlay" class="fixed inset-0 z-[100] flex items-center justify-center bg-[#1E4E4E]/95 backdrop-blur-md p-4">
    
    <div class="bg-white p-8 md:p-12 rounded-[8px] max-w-2xl w-full shadow-2xl relative border-t-[8px] border-[#6A9A9A]">
        
        <div class="flex justify-center mb-8">
            <div class="w-16 h-16 bg-[#F8FAFC] rounded-full flex items-center justify-center border-2 border-[#6A9A9A]">
                <span class="text-[#1E4E4E] font-['Montserrat'] font-bold text-xl">UCU</span>
            </div>
        </div>

        <h2 class="font-['Montserrat'] font-bold text-3xl text-[#1E4E4E] text-center mb-8 uppercase tracking-tight">
            Bienvenue chez <span class="text-[#6A9A9A]">UpcycleConnect</span>
        </h2>

        <div class="space-y-6 text-gray-700 mb-10 font-['Inter']">
            <div class="flex gap-5 items-start">
                <span class="flex-shrink-0 w-8 h-8 bg-[#6A9A9A]/10 text-[#6A9A9A] rounded-full flex items-center justify-center font-bold">1</span>
                <p class="text-base leading-relaxed">
                    <strong>Déposez vos objets :</strong> Accédez à nos box sécurisées grâce à un code-barres unique généré depuis votre espace.
                </p>
            </div>
            <div class="flex gap-5 items-start">
                <span class="flex-shrink-0 w-8 h-8 bg-[#6A9A9A]/10 text-[#6A9A9A] rounded-full flex items-center justify-center font-bold">2</span>
                <p class="text-base leading-relaxed">
                    <strong>Suivez votre impact :</strong> Visualisez votre <strong>Upcycling Score</strong> pour mesurer vos économies de ressources.
                </p>
            </div>
            <div class="flex gap-5 items-start">
                <span class="flex-shrink-0 w-8 h-8 bg-[#6A9A9A]/10 text-[#6A9A9A] rounded-full flex items-center justify-center font-bold">3</span>
                <p class="text-base leading-relaxed">
                    <strong>Participez aux ateliers :</strong> Retrouvez nos formateurs au <strong>174, rue La Fayette</strong> pour apprendre l'upcycling.
                </p>
            </div>
        </div>

        <div class="flex justify-center">
            <button onclick="document.getElementById('tuto-overlay').remove()" 
                    class="bg-[#1E4E4E] text-white px-10 py-4 rounded-[8px] font-['Montserrat'] font-bold uppercase tracking-widest hover:bg-[#6A9A9A] transition-all shadow-lg active:scale-95">
                Commencer l'expérience
            </button>
        </div>

        <p class="text-[10px] text-gray-400 text-center mt-10 font-['Inter'] italic">
            Conçu par Senard, Florian & Andi • Projet 2025-2026
        </p>
    </div>
</div>

    <nav class="bg-white border-b border-gray-100 sticky top-0 z-40">
        <div class="max-w-7xl mx-auto px-6 h-20 flex items-center justify-between">
            <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-ucu-foret rounded-ucu flex items-center justify-center text-white font-montserrat font-bold">UCU</div>
                <span class="font-montserrat text-xl text-ucu-foret tracking-tight">UpcycleConnect</span>
            </div>
            <div class="hidden md:flex gap-8 font-medium text-sm text-ucu-foret">
                <a href="#concept" class="hover:text-ucu-lagon transition">Concept</a>
                <a href="#services" class="hover:text-ucu-lagon transition">Prestations</a>
                <a href="#impact" class="hover:text-ucu-lagon transition">Impact</a>
            </div>
            <div class="flex items-center gap-4">
                <span class="text-xs font-bold text-gray-400">FR | EN</span>
                <a href="login/connexion.php" class="bg-ucu-foret text-white px-5 py-2 rounded-ucu text-sm font-bold hover:bg-ucu-lagon transition">Connexion</a>
            </div>
        </div>
    </nav>

    <header class="relative bg-white pt-16 pb-24 px-6 overflow-hidden">
        <div class="max-w-7xl mx-auto grid lg:grid-cols-2 gap-12 items-center">
            <div class="relative z-10">
                <div class="inline-flex items-center gap-2 px-3 py-1 bg-ucu-lagon/10 rounded-full mb-6">
                    <span class="w-2 h-2 bg-ucu-lagon rounded-full animate-pulse"></span>
                    <span class="text-xs font-bold text-ucu-lagon uppercase tracking-widest">Innovation Écologique • Depuis 2021</span>
                </div>
                <h1 class="font-montserrat text-5xl md:text-7xl text-ucu-foret leading-tight mb-6">
                    Donnez une <span class="text-ucu-lagon italic">seconde vie</span> à vos objets.
                </h1>
                <p class="text-lg text-gray-600 mb-10 max-w-lg leading-relaxed font-light">
                    La plateforme de référence en Europe pour l’upcycling intelligent. Relions particuliers et artisans pour transformer vos matériaux obsolètes en pièces d'exception.
                </p>
                <div class="flex flex-wrap gap-4">
                    <a href="#box" class="bg-ucu-foret text-white px-8 py-4 rounded-ucu font-bold shadow-xl shadow-ucu-foret/20 hover:-translate-y-1 transition">Déposer dans une Box</a>
                    <a href="#concept" class="border-2 border-ucu-lagon text-ucu-lagon px-8 py-4 rounded-ucu font-bold hover:bg-ucu-lagon hover:text-white transition">Découvrir le projet</a>
                </div>
            </div>
            <div class="relative">
                <div class="aspect-square bg-ucu-grey rounded-ucu border-2 border-dashed border-ucu-lagon/30 flex items-center justify-center overflow-hidden">
                    <span class="text-ucu-lagon/20 font-montserrat font-bold text-9xl">UCU</span>
                    </div>
                <div class="absolute -bottom-6 -left-6 bg-white p-4 rounded-ucu shadow-2xl border border-gray-100">
                    <p class="text-[10px] uppercase font-bold text-gray-400 mb-1 tracking-widest">Moyenne Communauté</p>
                    <p class="font-montserrat text-2xl text-ucu-foret font-bold">840 <span class="text-ucu-lagon font-normal text-sm">pts</span></p>
                </div>
            </div>
        </div>
    </header>

    <section id="services" class="py-24 bg-ucu-grey px-6">
        <div class="max-w-7xl mx-auto">
            <div class="flex flex-col md:flex-row justify-between items-end mb-16 gap-4">
                <div class="max-w-xl">
                    <h2 class="font-montserrat text-3xl text-ucu-foret mb-4 italic">Nos services de valorisation</h2>
                    <p class="text-gray-500">Un écosystème complet pour assurer le suivi complet de vos projets d'upcycling, de l'idée à la réalisation.</p>
                </div>
                <div class="h-1 w-24 bg-ucu-lagon rounded-full hidden md:block"></div>
            </div>

            <div class="grid md:grid-cols-3 gap-8">
                <div class="bg-white p-8 rounded-ucu shadow-sm hover:shadow-xl transition-all border-b-4 border-transparent hover:border-ucu-lagon group">
                    <div class="text-3xl mb-6">🔍</div>
                    <h3 class="font-montserrat font-bold text-lg text-ucu-foret mb-3 uppercase tracking-tight">Inventaire Intelligent</h3>
                    <p class="text-sm text-gray-600 leading-relaxed">Filtres avancés par type, localisation et potentiel de projet pour trouver les matériaux parfaits.</p>
                </div>

                <div class="bg-white p-8 rounded-ucu shadow-sm hover:shadow-xl transition-all border-b-4 border-transparent hover:border-ucu-lagon group">
                    <div class="text-3xl mb-6">📦</div>
                    <h3 class="font-montserrat font-bold text-lg text-ucu-foret mb-3 uppercase tracking-tight">Box & Conteneurs</h3>
                    <p class="text-sm text-gray-600 leading-relaxed">Déposez vos objets en box via un code-barres unique. Validation rigoureuse par notre service administratif.</p>
                </div>

                <div class="bg-white p-8 rounded-ucu shadow-sm hover:shadow-xl transition-all border-b-4 border-transparent hover:border-ucu-lagon group">
                    <div class="text-3xl mb-6">📈</div>
                    <h3 class="font-montserrat font-bold text-lg text-ucu-foret mb-3 uppercase tracking-tight">Upcycling Score</h3>
                    <p class="text-sm text-gray-600 leading-relaxed">Mesurez précisément la quantité de déchets évités et partagez votre impact avec la communauté.</p>
                </div>
            </div>
        </div>
    </section>

    <section class="py-20 bg-white px-6">
        <div class="max-w-7xl mx-auto">
            <div class="bg-ucu-foret rounded-ucu p-12 text-white flex flex-col md:flex-row items-center justify-between gap-12">
                <div class="max-w-md">
                    <h2 class="font-montserrat text-2xl mb-4">Retrouvez-nous au cœur du 10ème</h2>
                    <p class="text-ucu-lagon text-sm leading-relaxed mb-6 font-medium italic">
                        Bureaux administratifs et ateliers : 174, rue La Fayette, Paris.
                    </p>
                    <p class="text-xs opacity-70">
                        Également présent à : Ivry, Montreuil, Bourg-la-Reine et nos nouvelles annexes en province.
                    </p>
                </div>
                <div class="w-full md:w-auto">
                    <a href="https://www.google.com/maps" target="_blank" class="bg-white text-ucu-foret px-8 py-3 rounded-ucu font-bold inline-block hover:bg-ucu-lagon hover:text-white transition uppercase text-xs tracking-widest">Voir sur la carte</a>
                </div>
            </div>
        </div>
    </section>

    <footer class="py-12 px-6 border-t border-gray-100">
        <div class="max-w-7xl mx-auto flex flex-col md:flex-row justify-between items-center gap-8">
            <p class="text-xs text-gray-400 italic">© 2026 UpcycleConnect - Projet annuel 2025-2026. Réalisé par LACI S., GROLLEAU F., HASAJ A.</p>
            <div class="flex gap-6 grayscale opacity-50">
                <div class="w-5 h-5 bg-gray-400 rounded-full"></div>
                <div class="w-5 h-5 bg-gray-400 rounded-full"></div>
                <div class="w-5 h-5 bg-gray-400 rounded-full"></div>
            </div>
        </div>
    </footer>

</body>
</html>