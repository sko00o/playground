import sys
import soundfile
import torch

from pathlib import Path

current_dir = Path(__file__).parent
sys.path.insert(0, str(current_dir))
sys.path.insert(0, str(current_dir / "vits"))

import vits.commons as commons
import vits.utils as utils

from vits.models import SynthesizerTrn
from vits.text import text_to_sequence


def get_text(text, hps):
    text_norm = text_to_sequence(text, hps.data.text_cleaners)
    if hps.data.add_blank:
        text_norm = commons.intersperse(text_norm, 0)
    text_norm = torch.LongTensor(text_norm)
    return text_norm


class TTService:
    def __init__(self, cfg, model, speed):
        self.hps = utils.get_hparams_from_file(cfg)
        self.speed = speed

        symbols = self.hps.symbols
        import vits.text

        vits.text._symbol_to_id = {s: i for i, s in enumerate(symbols)}
        vits.text._id_to_symbol = {i: s for i, s in enumerate(symbols)}

        self.net_g = SynthesizerTrn(
            len(self.hps.symbols),
            self.hps.data.filter_length // 2 + 1,
            self.hps.train.segment_size // self.hps.data.hop_length,
            **self.hps.model
        ).cuda()
        _ = self.net_g.eval()
        _ = utils.load_checkpoint(model, self.net_g, None)

    def read(self, text):
        text = text.replace("~", "！")
        stn_tst = get_text(text, self.hps)
        with torch.no_grad():
            x_tst = stn_tst.cuda().unsqueeze(0)
            x_tst_lengths = torch.LongTensor([stn_tst.size(0)]).cuda()
            audio = (
                self.net_g.infer(
                    x_tst,
                    x_tst_lengths,
                    noise_scale=0.667,
                    noise_scale_w=0.2,
                    length_scale=self.speed,
                )[0][0, 0]
                .data.cpu()
                .float()
                .numpy()
            )
        return audio

    def read_save(self, text, filename):
        au = self.read(text)
        soundfile.write(filename, au, self.hps.data.sampling_rate)


def main():
    cfg = "models/config.json"
    model = "models/model.pth"
    speed = 1.0
    text = "你的中文怎么样，我觉得不是很好，但是我可以教你，哈哈哈"
    output_path = "output/test.wav"
    tt_service = TTService(cfg, model, speed)
    tt_service.read_save(text, output_path)


if __name__ == "__main__":
    main()
