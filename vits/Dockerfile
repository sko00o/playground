FROM pytorch/pytorch:2.6.0-cuda11.8-cudnn9-runtime

WORKDIR /workspace

RUN apt-get update && apt-get install -y git wget gcc g++ cmake libsndfile1

RUN git clone https://github.com/CjangCjengh/vits.git && cd vits && git checkout 43c74f369834294464eca69316048d2f9a7d54f1

WORKDIR /workspace/vits

COPY requirements.txt .
RUN pip install -r requirements.txt
RUN cd monotonic_align && \
    mkdir monotonic_align && \
    python3 setup.py build_ext --inplace

WORKDIR /workspace

# Chinese dialect lexicons that required by vits
RUN wget https://huggingface.co/spaces/skytnt/moe-tts/resolve/main/chinese_dialect_lexicons/jyutjyu.json
RUN wget https://huggingface.co/spaces/skytnt/moe-tts/resolve/main/chinese_dialect_lexicons/jyutjyu.ocd2
RUN wget https://huggingface.co/spaces/skytnt/moe-tts/resolve/main/chinese_dialect_lexicons/zaonhe.json
RUN wget https://huggingface.co/spaces/skytnt/moe-tts/resolve/main/chinese_dialect_lexicons/zaonhe.ocd2

COPY main.py .

CMD ["python3", "main.py"]
