FROM alpine AS runtime
COPY --from=builder /dist/surge /surge
RUN chmod +x /surge
CMD ["/surge"]
