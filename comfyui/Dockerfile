FROM pytorch/pytorch:2.4.0-cuda12.1-cudnn9-runtime

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && \
    apt-get install -yq git build-essential

ARG COMFYCLI=comfy-cli
RUN pip install --no-cache-dir ${COMFYCLI}

RUN comfy --skip-prompt --here install --cpu

RUN comfy set-default $(pwd)/ComfyUI --launch-extras="--listen --preview-method=auto"

EXPOSE 8188
CMD ["comfy", "launch"]