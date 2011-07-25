include $(GOROOT)/src/Make.inc

ALL=simple dynamic_pads

all: $(ALL)

clean:
	rm -rf *.[68] $(ALL)

%: %.go
	$(GC) $*.go
	$(LD) -o $@ $@.$O
