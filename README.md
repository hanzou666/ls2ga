# ls2ga

Augment reference genome by structural information from split-alignment

## usage

Local align contigs to reference sequence using last-split

```bash
lastdb -uNEAR -R01 db target.fa
lastal db query.fa | last-split > out.maf
```

Join split-alignments to a global alignment on a flat graph for each contig

```bash
ls2ga -M 32 -m 10000 -g flat.gfa -j aln.json target.fa query.fa out.maf
```

Transform graph and alignment in text format to in protobuf format

```bash
vg view -Fv flat.gfa > flat.vg
vg view -aJG aln.json > aln.gam
```

Create reference graph embedded differences in global-alignment

```bash
vg augment -i flat.vg aln.gam > aug.vg
```
