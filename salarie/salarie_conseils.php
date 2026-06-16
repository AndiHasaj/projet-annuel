<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>UCU Salarié | Rédaction de Conseils</title>
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

    <nav class="bg-ucu-lagon border-b border-white/10 sticky top-0 z-50 text-white">
        <div class="max-w-7xl mx-auto px-6 h-20 flex items-center justify-between">
            <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-white rounded-ucu flex items-center justify-center text-ucu-lagon font-montserrat font-bold italic">SAL</div>
                <h1 class="font-montserrat font-bold text-sm uppercase tracking-widest text-ucu-foret">Espace Salarié</h1>
            </div>
            <div class="flex gap-6 text-[10px] font-bold uppercase tracking-widest text-ucu-foret">
                <a href="salarie_dashboard.php" class="hover:text-white transition-colors">Dashboard</a>
                <a href="salarie_evenements.php" class="hover:text-white transition-colors">Événements</a>
                <a href="salarie_conseils.php" class="border-b-2 border-ucu-foret pb-1">Conseils</a>
            </div>
        </div>
    </nav>

    <main class="max-w-7xl mx-auto px-6 py-12">
        
        <div class="grid grid-cols-1 lg:grid-cols-12 gap-12">
            
            <section class="lg:col-span-8 space-y-8">
                <div class="bg-white p-8 rounded-ucu shadow-sm border border-gray-100">
                    <h2 class="font-montserrat font-bold text-xl text-ucu-foret uppercase mb-8 italic tracking-tight border-b pb-4">Rédiger un nouveau guide</h2>
                    
                    <form class="space-y-6">
                        <div>
                            <label class="text-[9px] font-bold uppercase text-gray-400 mb-2 block">Titre du conseil / News</label>
                            <input type="text" placeholder="ex: 5 astuces pour restaurer du cuir ancien" class="w-full bg-ucu-grey p-4 rounded-ucu text-xs outline-none focus:ring-2 focus:ring-ucu-lagon/20 font-semibold">
                        </div>

                        <div class="grid grid-cols-2 gap-6">
                            <div>
                                <label class="text-[9px] font-bold uppercase text-gray-400 mb-2 block">Catégorie</label>
                                <select class="w-full bg-ucu-grey p-4 rounded-ucu text-xs outline-none">
                                    <option>Techniques de base</option>
                                    <option>News Écologie</option>
                                    <option>Tutoriel Expert</option>
                                </select>
                            </div>
                            <div>
                                <label class="text-[9px] font-bold uppercase text-gray-400 mb-2 block">Mots-clés (Tags)</label>
                                <input type="text" placeholder="cuir, entretien, restauration" class="w-full bg-ucu-grey p-4 rounded-ucu text-xs outline-none">
                            </div>
                        </div>

                        <div>
                            <label class="text-[9px] font-bold uppercase text-gray-400 mb-2 block">Corps du texte (Markdown supporté)</label>
                            <div class="border border-gray-100 rounded-ucu overflow-hidden">
                                <div class="bg-ucu-grey p-2 border-b border-gray-100 flex gap-2">
                                    <button type="button" class="w-8 h-8 bg-white rounded border border-gray-200 font-bold text-xs">B</button>
                                    <button type="button" class="w-8 h-8 bg-white rounded border border-gray-200 italic text-xs">I</button>
                                    <button type="button" class="w-8 h-8 bg-white rounded border border-gray-200 text-xs">🔗</button>
                                </div>
                                <textarea rows="12" class="w-full p-4 text-xs outline-none resize-none" placeholder="Rédigez votre contenu technique ici..."></textarea>
                            </div>
                        </div>

                        <div class="flex items-center justify-between pt-6">
                            <button type="button" class="text-[9px] font-bold text-gray-400 uppercase tracking-widest hover:text-ucu-foret">Enregistrer en brouillon</button>
                            <button type="submit" class="bg-ucu-foret text-white px-8 py-4 rounded-ucu font-montserrat font-bold text-[10px] uppercase tracking-[0.2em] shadow-lg hover:bg-ucu-lagon transition-all">
                                Soumettre au responsable
                            </button>
                        </div>
                    </form>
                </div>
            </section>

            <aside class="lg:col-span-4 space-y-8">
                <div class="bg-white p-6 rounded-ucu shadow-sm border border-gray-100">
                    <h3 class="font-montserrat font-bold text-[10px] uppercase text-ucu-foret mb-4 tracking-widest italic">Image de couverture</h3>
                    <div class="aspect-video bg-ucu-grey rounded-ucu border-2 border-dashed border-gray-200 flex flex-col items-center justify-center cursor-pointer hover:border-ucu-lagon transition-all">
                        <span class="text-2xl mb-2">🖼️</span>
                        <p class="text-[8px] font-bold uppercase text-gray-400">Glisser-déposer une photo</p>
                    </div>
                </div>

                <div class="bg-white rounded-ucu shadow-sm border border-gray-100 overflow-hidden">
                    <div class="p-4 border-b border-gray-100 bg-ucu-grey/30 text-center">
                        <h3 class="font-montserrat font-bold text-[9px] uppercase text-ucu-foret tracking-widest">Mes publications</h3>
                    </div>
                    <div class="divide-y divide-gray-50">
                        <div class="p-4 hover:bg-ucu-grey transition-all">
                            <p class="text-[10px] font-bold text-ucu-foret uppercase">Le guide du ponçage</p>
                            <div class="flex justify-between items-center mt-2">
                                <span class="text-[8px] text-green-500 font-bold uppercase">En ligne</span>
                                <span class="text-[8px] text-gray-400 font-mono">1.2k vues</span>
                            </div>
                        </div>
                        <div class="p-4 hover:bg-ucu-grey transition-all">
                            <p class="text-[10px] font-bold text-gray-400 uppercase italic">Zéro Déchet au bureau</p>
                            <div class="flex justify-between items-center mt-2">
                                <span class="text-[8px] text-orange-400 font-bold uppercase">En révision</span>
                                <span class="text-[8px] text-gray-400">---</span>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="bg-ucu-lagon/10 p-6 rounded-ucu border-l-4 border-ucu-lagon">
                    <p class="text-[9px] font-bold text-ucu-lagon uppercase mb-2">📌 Rappel Charte</p>
                    <p class="text-[10px] text-ucu-foret italic leading-relaxed">
                        N'oubliez pas d'inclure des étapes claires et des conseils sécurité pour que les particuliers augmentent leur <strong>Upcycling Score</strong> sans risque.
                    </p>
                </div>
            </aside>

        </div>
    </main>

    <footer class="mt-20 py-10 border-t border-gray-100 text-center">
        <p class="text-[9px] font-bold text-gray-400 uppercase tracking-[0.3em]">UpcycleConnect Salarié • Rédaction & News • 2026</p>
    </footer>

</body>
</html>