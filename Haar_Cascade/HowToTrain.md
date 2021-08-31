# Haar training

### Reference

[https://m.blog.naver.com/PostView.nhn?isHttpsRedirect=true&blogId=hihimani&logNo=80154867344&proxyReferer=](https://m.blog.naver.com/PostView.nhn?isHttpsRedirect=true&blogId=hihimani&logNo=80154867344&proxyReferer=)

[https://velog.io/@huttzza/실시간-얼굴-인식-프로그램-1차-구현](https://velog.io/@huttzza/%EC%8B%A4%EC%8B%9C%EA%B0%84-%EC%96%BC%EA%B5%B4-%EC%9D%B8%EC%8B%9D-%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%A8-1%EC%B0%A8-%EA%B5%AC%ED%98%84)

[https://darkpgmr.tistory.com/70?category=460965](https://darkpgmr.tistory.com/70?category=460965)

## Contents

- OpenCV installation
- How to train Haar

## OpenCV installation

[https://webnautes.tistory.com/1186](https://webnautes.tistory.com/1186)

## How to train Haar

- Step 1 : Collect Pupil Images
    1. run vid2frame.ipynb

- Step 2 : Set Pupil ROI & Divide into positive and negative data

    ***window 10*** 

    1. set pupil ROI using maker program
    2. run img2gray.ipynb
    3. run img2PnN.ipynb

- Step 3 : Train Haar by LBP method

    ***Ubuntu 18.04***

    1. clean_all.sh
    2. make_txt.sh
    3. 1_Train.sh
    4. 2_Train.sh
    5. 3_Train.sh
    6. 4_Train.sh 

    *** `struct.error: unpack requires a string argument of length 12` ***

    [https://freesoft.dev/program/158147598](https://freesoft.dev/program/158147598)

    Check your .vec files. If the size of file is zero byte, delete them. 

---

by The Legendary Dragon 🐉