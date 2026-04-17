package entities;
import observer.MovieSubject;

public class Movie extends MovieSubject {
    private String id;
    private String title;
    private String description;
    private int duration;

    public Movie(String id, String title, String description, int duration) {
        this.id = id;
        this.title = title;
        this.description = description;
        this.duration = duration;
    }

    public String getId() { return id; }
    public String getTitle() { return title; }
}
