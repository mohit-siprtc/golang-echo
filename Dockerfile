# Use the official PostgreSQL image
FROM postgres:15

# Set environment variables for the database
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=admin
ENV POSTGRES_DB=book_admins

# Copy custom configuration files (if any)
# Uncomment the following line if you have a custom `postgresql.conf` file
# COPY postgresql.conf /etc/postgresql/postgresql.conf

# Expose the PostgreSQL default port
EXPOSE 5433

# Default command (already defined in the base image)
CMD ["postgres"]
