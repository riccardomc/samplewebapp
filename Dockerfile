FROM scratch
EXPOSE 8080
ADD samplewebapp /
CMD ["/samplewebapp"]
