using CSV, DataFrames, ChemometricsTools
deef = CSV.File("/home/caseykneale/Desktop/langset.csv") |> DataFrame

l = deef[:,1]
langs = [:C,Symbol("C++"),:Go,:Rust,:FORTRAN,:Java,:Scala,:Julia,:Python,:Haskell,:JavaScript,Symbol("F#"),
            :VB,:Excel,:QBASIC,:R,:Matlab,:Mathematica,Symbol("C#")]
data = deef[:,2:end] |> Matrix .|> Float64  

cs = CenterScale(data)
pca = PCA(cs(data))

using Plots

scatter(pca.Scores[:,1],pca.Scores[:,2], color = :white,markerstrokewidth=0)
annotate!(pca.Scores[:,1],pca.Scores[:,2],langs, legend = false)
png("langs.png")
