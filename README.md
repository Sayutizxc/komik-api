# Komik Api

- <a href="#daftar-komik-terbaru">Daftar Komik terbaru<a/>
- <a href="#detail-komik">Detail Komik<a/>
- <a href="#chapter">Chapter<a/>
- <a href="#pencarian-komik">Pencarian Komik<a/>

## Daftar Komik Terbaru

```http
  https://api.sayuti.codes/comic/home/{page}?limit={limit}
  https://api.sayuti.codes/comic/home/
  https://api.sayuti.codes/comic/home/1?limit=20
```

Parameter page bersifat opsional,defaultnya adalah 1, yaitu halaman pertama, parameter page hanya dapat di isi dengan angka,dan harus lebih dari 0

Query parameter limit bersifat opsional, defaultnya adalah 32,dan harus lebih dari 0, berfungsi untuk membatasi berapa banyak data yang akan didapat dalam satu halamannya

### Contoh

> https://api.sayuti.codes/comic/home/1?limit=2

```json
  {
  "status": 200,
  "data": [
    {
      "url": "https://klikmanga.com/manga/yuan-zun/",
      "title": "Yuan Zun",
      "thumbnail": "https://klikmanga.com/wp-content/uploads/2018/08/23044-175x238.jpg",
      "score": "4.7",
      "LastChapter": [
        {
          "chapter_url": "https://klikmanga.com/manga/yuan-zun/chapter-258/",
          "chapter": "Chapter 258",
          "date": "10 jam yang lalu"
        },
        {
          "chapter_url": "https://klikmanga.com/manga/yuan-zun/chapter-257-5/",
          "chapter": "Chapter 257.5",
          "date": "10 jam yang lalu"
        }
      ]
    },
    {
      "url": "https://klikmanga.com/manga/yofukashi-no-uta/",
      "title": "Yofukashi no Uta",
      "thumbnail": "https://klikmanga.com/wp-content/uploads/2019/11/40969-1-175x238.jpg",
      "score": "4.7",
      "LastChapter": [
        {
          "chapter_url": "https://klikmanga.com/manga/yofukashi-no-uta/chapter-68/",
          "chapter": "Chapter 68",
          "date": "10 jam yang lalu"
        },
        {
          "chapter_url": "https://klikmanga.com/manga/yofukashi-no-uta/chapter-67/",
          "chapter": "Chapter 67",
          "date": "Okt 15, 2021"
        }
      ]
    }
  ]
}
```


## Detail Komik

```http
  https://api.sayuti.codes/comic/detail?url={url}
  https://api.sayuti.codes/comic/detail?url=https://klikmanga.com/manga/yuan-zun/
```

Query parameter url wajib di isi dengan url dari komik yang ingin dilihat detailnya

### Contoh

> https://api.sayuti.codes/comic/detail?url=https://klikmanga.com/manga/isekai-majutsushi-wa-mahou-wo-tonaenai/

```json
  {
  "status": 200,
  "data": {
    "url": "https://klikmanga.com/manga/isekai-majutsushi-wa-mahou-wo-tonaenai/",
    "title": "Isekai Majutsushi wa Mahou wo Tonaenai",
    "thumbnail": "https://klikmanga.com/wp-content/uploads/2020/06/50056-193x278.jpeg",
    "rating": "4.8",
    "authors": "Mochi",
    "artists": "Koppamu",
    "genres": [
      "Adventure",
      "Ecchi",
      "Fantasy",
      "Harem",
      "Isekai",
      "Mature",
      "Romance",
      "Seinen"
    ],
    "type": "Manga",
    "release": "2020",
    "status": "OnGoing",
    "synopsis": "Penyihir bernama Yard tiba tiba dipanggil ke dunia yang berbeda sebagai seorang Pahlawan. Penyihir di dunia ini tak sebanding jika dibandingkan dengan Yard yang datang dari dunia lain. Pahlawan lain yang dipanggil juga tak bisa di andalkan. Lebih dari itu, tidak mungkin kembali ke dunia yang lama lagi… Di dalam keputusasaan tersebut, Yard berusaha seorang diri. Bersiaplah, manusia di dunia ini. Dunia dimana ilmu sihir tertinggal. Aku punya mainan yang luar biasa di dalam box ini. Ijinkan aku bermain dengannya. Penyihir jahat bernama Yard menaklukkan dunia lain dengan Sihir Hitam yang luar biasa!",
    "chapters": [
      {
        "chapter_url": "https://klikmanga.com/manga/isekai-majutsushi-wa-mahou-wo-tonaenai/chapter-8/",
        "chapter": "Chapter 8",
        "date": "20 jam yang lalu"
      },
      {
        "chapter_url": "https://klikmanga.com/manga/isekai-majutsushi-wa-mahou-wo-tonaenai/chapter-7/",
        "chapter": "Chapter 7",
        "date": "Jan 22, 2021"
      },
      {
        "chapter_url": "https://klikmanga.com/manga/isekai-majutsushi-wa-mahou-wo-tonaenai/chapter-6/",
        "chapter": "Chapter 6",
        "date": "Jan 16, 2021"
      },
      {
        "chapter_url": "https://klikmanga.com/manga/isekai-majutsushi-wa-mahou-wo-tonaenai/chapter-5/",
        "chapter": "Chapter 5",
        "date": "Nov 19, 2020"
      },
      {
        "chapter_url": "https://klikmanga.com/manga/isekai-majutsushi-wa-mahou-wo-tonaenai/chapter-4/",
        "chapter": "Chapter 4",
        "date": "Okt 31, 2020"
      },
      {
        "chapter_url": "https://klikmanga.com/manga/isekai-majutsushi-wa-mahou-wo-tonaenai/chapter-3/",
        "chapter": "Chapter 3",
        "date": "Okt 26, 2020"
      },
      {
        "chapter_url": "https://klikmanga.com/manga/isekai-majutsushi-wa-mahou-wo-tonaenai/chapter-2/",
        "chapter": "Chapter 2",
        "date": "Jul 27, 2020"
      },
      {
        "chapter_url": "https://klikmanga.com/manga/isekai-majutsushi-wa-mahou-wo-tonaenai/chapter-1/",
        "chapter": "Chapter 1",
        "date": "Jun 15, 2020"
      }
    ],
    "comments": [
      {
        "name": "no coment",
        "profile_pic": "https://secure.gravatar.com/avatar/958ae624b5eaa7e012c2107e5f18e69b?s=60&d=mm&r=g",
        "comment": "Up up up up up up",
        "date": "Apr 2, 2021 at 4:19 pm"
      },
      {
        "name": "Leevhonk",
        "profile_pic": "https://secure.gravatar.com/avatar/cc8f1e1c1f90acdd8bc815c9a3f5e4a3?s=60&d=mm&r=g",
        "comment": "Komik macam apa ini…!! Bahhh Lanjutkan..",
        "date": "Okt 2, 2020 at 12:56 am"
      },
      {
        "name": "Sya-kun",
        "profile_pic": "https://klikmanga.com/wp-content/uploads/2020/08/IMG_20200708_110721_110-1.jpg",
        "comment": "Lanjutkan,min… Ditunggu chapter selanjutnya",
        "date": "Jun 18, 2020 at 11:16 am"
      },
      {
        "name": "Albaitu",
        "profile_pic": "https://secure.gravatar.com/avatar/5f9b3896067817113493fc0d7565a2e5?s=60&d=mm&r=g",
        "comment": "Tiap hari min update",
        "date": "Jun 15, 2020 at 8:00 pm"
      },
      {
        "name": "Aditama",
        "profile_pic": "https://secure.gravatar.com/avatar/a8bc86b911042073ca0bb483bb88a140?s=60&d=mm&r=g",
        "comment": "Lanjut min",
        "date": "Jun 15, 2020 at 7:15 pm"
      },
      {
        "name": "karelhasegawa",
        "profile_pic": "https://secure.gravatar.com/avatar/3e970efb3b28b3432850684f5541551f?s=60&d=mm&r=g",
        "comment": "lanjut kan min penasaran kedepan nya huahahahahaha",
        "date": "Jun 15, 2020 at 1:33 pm"
      }
    ]
  }
}
```

## Chapter

```http
  https://api.sayuti.codes/comic/chapter?url={url}
  https://api.sayuti.codes/comic/chapter?url=https://klikmanga.com/manga/isekai-majutsushi-wa-mahou-wo-tonaenai/chapter-1/
```

Query parameter url wajib di isi dengan url dari chapter yang ingin dilihat gambarnya

### Contoh

> https://api.sayuti.codes/comic/chapter?url=https://klikmanga.com/manga/isekai-majutsushi-wa-mahou-wo-tonaenai/chapter-2/

```json
  {
  "status": 200,
  "data": {
    "chapter": "Chapter 2",
    "url": "https://klikmanga.com/manga/isekai-majutsushi-wa-mahou-wo-tonaenai/chapter-2/",
    "next": "https://klikmanga.com/manga/isekai-majutsushi-wa-mahou-wo-tonaenai/chapter-3/?style=list",
    "prev": "https://klikmanga.com/manga/isekai-majutsushi-wa-mahou-wo-tonaenai/chapter-1/?style=list",
    "images": [
      "https://i3.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/001.jpg",
      "https://i3.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/002.jpg",
      "https://i1.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/003.jpg",
      "https://i0.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/004.jpg",
      "https://i2.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/005.jpg",
      "https://i3.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/006.jpg",
      "https://i2.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/007.jpg",
      "https://i3.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/008.jpg",
      "https://i3.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/009.jpg",
      "https://i0.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/010.jpg",
      "https://i0.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/011.jpg",
      "https://i3.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/012.jpg",
      "https://i1.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/013.jpg",
      "https://i2.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/014.jpg",
      "https://i1.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/015.jpg",
      "https://i2.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/016.jpg",
      "https://i1.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/017.jpg",
      "https://i0.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/018.jpg",
      "https://i1.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/019.jpg",
      "https://i0.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/020.jpg",
      "https://i2.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/021.jpg",
      "https://i0.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/022.jpg",
      "https://i2.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/023.jpg",
      "https://i1.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/024.jpg",
      "https://i1.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/025.jpg",
      "https://i1.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/026.jpg",
      "https://i2.wp.com/img.klikmanga.com/data/manga_5ee697df588a9/chapter-2/027.jpg"
    ],
    "Comments": [
      {
        "name": "Zeraldoo",
        "profile_pic": "https://secure.gravatar.com/avatar/53fd4087ebb76cbcaa47328bd064c52f?s=60&d=mm&r=g",
        "comment": "Up",
        "date": "Sep 13, 2020 at 4:31 pm"
      }
    ]
  }
}
```

## Pencarian Komik

```http
  https://api.sayuti.codes/comic/search/{page}?s={keyword}&limit={limit}
  https://api.sayuti.codes/comic/search/1?s=dark&limit=2
  https://api.sayuti.codes/comic/search/?s=dark
```

Parameter page bersifat opsional,defaultnya adalah 1, yaitu halaman pertama, parameter page hanya dapat di isi dengan angka,dan harus lebih dari 0

Query parameter s wajib di isi, ini berisi dengan keyword yang ingin dicari

Query parameter limit bersifat opsional, defaultnya adalah 32,dan harus lebih dari 0, berfungsi untuk membatasi berapa banyak data yang akan didapat dalam satu halamannya

### Contoh

> https://api.sayuti.codes/comic/search/1?s=dark&limit=2

```json
  {
  "status": 200,
  "data": [
    {
      "url": "https://klikmanga.com/manga/dark-hole-another-survivor/",
      "title": "Dark Hole: Another Survivor",
      "thumbnail": "https://klikmanga.com/wp-content/uploads/2021/10/Dark-Hole-Another-Survivor-193x278.jpg",
      "score": "0",
      "LastChapter": [
        {
          "chapter_url": "https://klikmanga.com/manga/dark-hole-another-survivor/chapter-1/",
          "chapter": "Chapter 1",
          "date": "12 jam yang lalu"
        }
      ]
    },
    {
      "url": "https://klikmanga.com/manga/chotto-dake-ai-ga-omoi-dark-elf-ga-isekai-kara-oikakete-kita/",
      "title": "Chotto Dake ai ga Omoi Dark Elf ga Isekai Kara Oikakete Kita",
      "thumbnail": "https://klikmanga.com/wp-content/uploads/2021/09/Chotto-Dake-ai-ga-Omoi-Dark-Elf-ga-Isekai-Kara-Oikakete-Kita-193x278.jpeg",
      "score": "5",
      "LastChapter": [
        {
          "chapter_url": "https://klikmanga.com/manga/chotto-dake-ai-ga-omoi-dark-elf-ga-isekai-kara-oikakete-kita/chapter-1/",
          "chapter": "Chapter 1",
          "date": "2021-09-21 18:08:38"
        }
      ]
    }
  ]
}
```


