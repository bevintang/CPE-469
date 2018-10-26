CC = mpicc
CFLAGS = -g -Wall

seq: heatTransfer.c
	$(CC) $(CFLAGS) -o heatC.out heatTransfer.c

dis: heatTransferMPI.c
	$(CC) $(CFLAGS) -o heatMPI.out heatTransferMPI.c

run:
	mpirun -hostfile hostfile -np 4 heatMPI.out

clean:
	rm *.out