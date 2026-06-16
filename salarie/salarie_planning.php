<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>UCU Salarié | Mon Planning</title>
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
                <a href="salarie_planning.php" class="border-b-2 border-ucu-foret pb-1">Planning</a>
                <a href="salarie_moderation.php" class="hover:text-white transition-colors">Forums</a>
            </div>
        </div>
    </nav>

    <main class="max-w-7xl mx-auto px-6 py-12">
        
        <div class="flex flex-col md:flex-row justify-between items-end mb-10 gap-4">
            <div>
                <h2 class="font-montserrat font-bold text-3xl text-ucu-foret uppercase tracking-tighter">Mon <span class="text-ucu-lagon italic">Planning</span></h2>
                <p class="text-sm text-gray-400 mt-2">Semaine du 16 au 22 Mars 2026</p>
            </div>
            <div class="flex gap-2">
                <button class="bg-white border border-gray-200 p-3 rounded-ucu hover:bg-ucu-grey">◀</button>
                <button class="bg-white border border-gray-200 p-3 rounded-ucu hover:bg-ucu-grey">▶</button>
                <button class="bg-ucu-foret text-white px-6 py-3 rounded-ucu font-bold text-[10px] uppercase tracking-widest">Aujourd'hui</button>
            </div>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-7 gap-4">
            
            <div class="hidden lg:contents">
                <div class="text-center p-4 font-montserrat font-bold text-[10px] uppercase text-gray-400">Lun. 16</div>
                <div class="text-center p-4 font-montserrat font-bold text-[10px] uppercase text-gray-400">Mar. 17</div>
                <div class="text-center p-4 font-montserrat font-bold text-[10px] uppercase text-ucu-lagon italic">Mer. 18</div>
                <div class="text-center p-4 font-montserrat font-bold text-[10px] uppercase text-gray-400">Jeu. 19</div>
                <div class="text-center p-4 font-montserrat font-bold text-[10px] uppercase text-gray-400">Ven. 20</div>
                <div class="text-center p-4 font-montserrat font-bold text-[10px] uppercase text-gray-400">Sam. 21</div>
                <div class="text-center p-4 font-montserrat font-bold text-[10px] uppercase text-gray-400">Dim. 22</div>
            </div>

            <div class="bg-white min-h-[400px] rounded-ucu border border-gray-100 p-2 space-y-2"></div>

            <div class="bg-white min-h-[400px] rounded-ucu border border-gray-100 p-2 space-y-2">
                <div class="bg-ucu-lagon/10 border-l-4 border-ucu-lagon p-3 rounded">
                    <p class="text-[9px] font-bold text-ucu-lagon uppercase">14h - 17h</p>
                    <p class="text-[10px] font-bold text-ucu-foret leading-tight mt-1">Formation Soudure Arc</p>
                    <p class="text-[8px] text-gray-400 italic">Atelier Central</p>
                </div>
            </div>

            <div class="bg-ucu-lagon/5 min-h-[400px] rounded-ucu border-2 border-ucu-lagon p-2 space-y-2 relative">
                <div class="bg-ucu-foret p-3 rounded shadow-md text-white">
                    <p class="text-[9px] font-bold text-ucu-lagon uppercase">09h - 12h</p>
                    <p class="text-[10px] font-bold leading-tight mt-1 uppercase">Réunion Équipe</p>
                    <p class="text-[8px] opacity-70 italic">Bureau Ivry</p>
                </div>
                <div class="bg-white border border-gray-100 p-3 rounded">
                    <p class="text-[9px] font-bold text-ucu-lagon uppercase">14h - 16h</p>
                    <p class="text-[10px] font-bold text-ucu-foret leading-tight mt-1">Modération Forums</p>
                    <p class="text-[8px] text-gray-400 italic">Distanciel</p>
                </div>
            </div>

            <div class="bg-white min-h-[400px] rounded-ucu border border-gray-100 p-2 space-y-2">
                <div class="bg-orange-50 border-l-4 border-orange-400 p-3 rounded">
                    <p class="text-[9px] font-bold text-orange-500 uppercase">10h - 12h</p>
                    <p class="text-[10px] font-bold text-ucu-foret leading-tight mt-1">Tuto : Vernis Bio</p>
                    <p class="text-[8px] text-orange-400 italic font-bold uppercase">En attente de validation</p>
                </div>
            </div>

            <div class="bg-white min-h-[400px] rounded-ucu border border-gray-100 p-2 space-y-2"></div>
            <div class="bg-white min-h-[400px] rounded-ucu border border-gray-100 p-2 space-y-2"></div>
            <div class="bg-white min-h-[400px] rounded-ucu border border-gray-100 p-2 space-y-2"></div>

        </div>

        <section class="mt-12 grid grid-cols-1 md:grid-cols-3 gap-6">
            <div class="bg-white p-6 rounded-ucu shadow-sm border border-gray-100 flex items-center justify-between">
                <div>
                    <p class="text-[9px] font-bold text-gray-400 uppercase">Heures cette semaine</p>
                    <p class="text-xl font-montserrat font-bold text-ucu-foret">24h / 35h</p>
                </div>
                <div class="w-12 h-12 bg-ucu-grey rounded-full flex items-center justify-center">🕒</div>
            </div>
            <div class="bg-white p-6 rounded-ucu shadow-sm border border-gray-100 flex items-center justify-between">
                <div>
                    <p class="text-[9px] font-bold text-gray-400 uppercase">Ateliers animés</p>
                    <p class="text-xl font-montserrat font-bold text-ucu-foret">4 sessions</p>
                </div>
                <div class="w-12 h-12 bg-ucu-grey rounded-full flex items-center justify-center">🎓</div>
            </div>
            <div class="bg-white p-6 rounded-ucu shadow-sm border border-gray-100 flex items-center justify-between">
                <div>
                    <p class="text-[9px] font-bold text-gray-400 uppercase">Score Formateur</p>
                    <p class="text-xl font-montserrat font-bold text-ucu-lagon">4.8 / 5</p>
                </div>
                <div class="w-12 h-12 bg-ucu-grey rounded-full flex items-center justify-center">⭐️</div>
            </div>
        </section>

    </main>

    <footer class="mt-20 py-10 border-t border-gray-100 text-center">
        <p class="text-[9px] font-bold text-gray-400 uppercase tracking-[0.3em]">UpcycleConnect Salarié • Planning Management • 2026</p>
    </footer>

</body>
</html>