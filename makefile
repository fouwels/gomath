all: bifurcation julia mandlebrot lorenz

bifurcation:
	go test -v -run TestBifurcation

julia:
	go test -v -run TestJulia

mandlebrot:
	go test -v -run TestMandlebrot
	
lorenz:
	go test -v -run TestLorenz

clean:
	rm -rf *.png
 