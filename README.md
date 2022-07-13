# Introduction

A simple API built to fetch chapters along with invocations from known book "Hisnul Muslim" (Fortress of the Muslim)

## Use Cases
### List chapters with their invocations
```http
GET /api/v1/invocations
```
### Response
```json
[
  {
    "id": 1,
    "chapter_name": "1.Çfarë duhet thënë kur të zgjohemi nga gjumi",
    "invocations": [
      {
        "id": 1,
        "audio": "/audios/001_01.mp3",
        "arabic": "الْحَمْدُ للهِ الَّذِي أَحْيَانَا بَعْدَ مَا أَمَاتَنَا وَإِلَيْهِ النُّشُورُ",
        "latin": "Elhamdulil-lahil-ledhi ahjana ba’de ma ematena, ve ilejhin-nushur",
        "albanian": "Falënderimi i takon Allahut, i Cili na ringjalli pasi na bëri të vdekur dhe tek Ai është ringjallja (tubimi Ditën e Gjykimit).",
        "reference": "Buhariu “Fet’hul-Bari” 11/113 dhe Muslimi 4/2083"
      },
      ...
```

### List chapters
```http
GET /api/v1/chapters
```
### Response
```json
[
  {
    "id": 1,
    "chapter_name": "1.Çfarë duhet thënë kur të zgjohemi nga gjumi",
  },
  {
    "id": 2,
    "chapter_name": "2.Duaja kur të veshim rrobat",
  }
      ...
```

### Single chapter
```http
GET /api/v1/chapters/:id
```
### Response
```json
{
    "id": 1,
    "chapter_name": "1.Çfarë duhet thënë kur të zgjohemi nga gjumi",
    "invocations": [
      {
        "id": 1,
        "audio": "/audios/001_01.mp3",
        "arabic": "الْحَمْدُ للهِ الَّذِي أَحْيَانَا بَعْدَ مَا أَمَاتَنَا وَإِلَيْهِ النُّشُورُ",
        "latin": "Elhamdulil-lahil-ledhi ahjana ba’de ma ematena, ve ilejhin-nushur",
        "albanian": "Falënderimi i takon Allahut, i Cili na ringjalli pasi na bëri të vdekur dhe tek Ai është ringjallja (tubimi Ditën e Gjykimit).",
        "reference": "Buhariu “Fet’hul-Bari” 11/113 dhe Muslimi 4/2083"
      },
      {
        "id": 2,
        "audio": "/audios/001_02.mp3",
        "arabic": "لَا إِلَهَ إِلَّا اللهُ وَحْدَهُ لَا شَرِيكَ لَهُ، لَهُ الْمُلْكُ وَلَهُ الْحَمْدَ، وَهُوَ عَلَى كُلِّ شَيْءٍ قَدِيرٌ. سُبْحَانَ اللهِ، وَالْحَمْدُ للهِ، ولَا إِلَهَ إِلَّا\\n\\nاللهُ، وَاللهُ أَكْبَرُ، وَلَا حَوْلَ وَلَا قُوَّةَ إِلَّا بِاللهِ العَلِيِّ الْعَظيِمِ، ربِّ اغْفِرلِي",
        "latin": "La ilahe il-lAll-llahu vahdehu la sherike leh, lehul-mulku ve lehul-hamdu ve huve ala kul-li shej’in kadir. SubhanAll-llah, vel-hamdulil-lah, ve la ilahe il-laAll-llah, vAll-llahu Ekber, ve la havle ve la kuvvete il-la bil-lahil Alij-jil Adhim. Rabbigfir li.",
        "albanian": "Nuk ka hyjni që meriton të adhurohet përveç Allahut, të Vetëm e i pa rival. Atij i takon Sundimi dhe Lavdërimi dhe Ai është i Gjithëfuqishëm mbi çdo gjë. I Madhërishëm qoftë Allahu, Falënderimi i përket vetëm Allahut, nuk ka hyjni që meriton të adhurohet përveç Allahut, Allahu është më i Madhi, nuk ka ndryshim e as forcë pa ndihmën e Allahut të Lartmadhëruar, Zoti im më fal.",
        "reference": "“Kush e thotë këtë (lutje) do t’i falen mëkatet, e nëse lutet do t’i pranohet lutja, e nëse ngritët e merr abdes dhe falet do t’i pranohet namazi.” Shënon Buhariu."
      },
      ...
```

### List specific chapter invocations
```http
GET /api/v1/chapters/:id/invocations
```
### Response
```json
[
      {
        "id": 1,
        "audio": "/audios/001_01.mp3",
        "arabic": "الْحَمْدُ للهِ الَّذِي أَحْيَانَا بَعْدَ مَا أَمَاتَنَا وَإِلَيْهِ النُّشُورُ",
        "latin": "Elhamdulil-lahil-ledhi ahjana ba’de ma ematena, ve ilejhin-nushur",
        "albanian": "Falënderimi i takon Allahut, i Cili na ringjalli pasi na bëri të vdekur dhe tek Ai është ringjallja (tubimi Ditën e Gjykimit).",
        "reference": "Buhariu “Fet’hul-Bari” 11/113 dhe Muslimi 4/2083"
      },
      {
        "id": 2,
        "audio": "/audios/001_02.mp3",
        "arabic": "لَا إِلَهَ إِلَّا اللهُ وَحْدَهُ لَا شَرِيكَ لَهُ، لَهُ الْمُلْكُ وَلَهُ الْحَمْدَ، وَهُوَ عَلَى كُلِّ شَيْءٍ قَدِيرٌ. سُبْحَانَ اللهِ، وَالْحَمْدُ للهِ، ولَا إِلَهَ إِلَّا\\n\\nاللهُ، وَاللهُ أَكْبَرُ، وَلَا حَوْلَ وَلَا قُوَّةَ إِلَّا بِاللهِ العَلِيِّ الْعَظيِمِ، ربِّ اغْفِرلِي",
        "latin": "La ilahe il-lAll-llahu vahdehu la sherike leh, lehul-mulku ve lehul-hamdu ve huve ala kul-li shej’in kadir. SubhanAll-llah, vel-hamdulil-lah, ve la ilahe il-laAll-llah, vAll-llahu Ekber, ve la havle ve la kuvvete il-la bil-lahil Alij-jil Adhim. Rabbigfir li.",
        "albanian": "Nuk ka hyjni që meriton të adhurohet përveç Allahut, të Vetëm e i pa rival. Atij i takon Sundimi dhe Lavdërimi dhe Ai është i Gjithëfuqishëm mbi çdo gjë. I Madhërishëm qoftë Allahu, Falënderimi i përket vetëm Allahut, nuk ka hyjni që meriton të adhurohet përveç Allahut, Allahu është më i Madhi, nuk ka ndryshim e as forcë pa ndihmën e Allahut të Lartmadhëruar, Zoti im më fal.",
        "reference": "“Kush e thotë këtë (lutje) do t’i falen mëkatet, e nëse lutet do t’i pranohet lutja, e nëse ngritët e merr abdes dhe falet do t’i pranohet namazi.” Shënon Buhariu."
      },
      ...
```

### Single chapter invocation
```http
GET /api/v1/chapters/:id/invocations/:invocationId
```
### Response
```json
{
        "id": 1,
        "audio": "/audios/001_01.mp3",
        "arabic": "الْحَمْدُ للهِ الَّذِي أَحْيَانَا بَعْدَ مَا أَمَاتَنَا وَإِلَيْهِ النُّشُورُ",
        "latin": "Elhamdulil-lahil-ledhi ahjana ba’de ma ematena, ve ilejhin-nushur",
        "albanian": "Falënderimi i takon Allahut, i Cili na ringjalli pasi na bëri të vdekur dhe tek Ai është ringjallja (tubimi Ditën e Gjykimit).",
        "reference": "Buhariu “Fet’hul-Bari” 11/113 dhe Muslimi 4/2083"
},
```

### Audio
```http
GET /audios/001_01.mp3
```
### Response
In this case a static mp3 file will be served
