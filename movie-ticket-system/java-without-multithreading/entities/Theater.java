package entities;
import java.util.ArrayList;
import java.util.List;

public class Theater {
    private String id;
    private String name;
    private String location;
    private List<Show> shows;

    public Theater(String id, String name, String location) {
        this.id = id;
        this.name = name;
        this.location = location;
        this.shows = new ArrayList<>();
    }

    public String getId() { return id; }
    public void addShow(Show show) { this.shows.add(show); }
}
